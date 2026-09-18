---
subcategory: "Unity Catalog"
---
# databricks_ai_gateway_model_services Data Source
[![GA](https://img.shields.io/badge/Release_Stage-GA-green)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/aigateway)

Lists the Unity Catalog model services that are visible to the current principal in a schema.


## Example Usage
The following example lists model services in the `main.default` schema:

```hcl
data "databricks_ai_gateway_model_services" "all" {
  parent = "schemas/main.default"
}

output "model_services" {
  value = data.databricks_ai_gateway_model_services.all.model_services
}
```


## Arguments
The following arguments are supported:
* `page_size` (integer, optional) - Maximum number of model services to return. Defaults to 100 when unset or 0;
  the maximum is 100. Use `page_token` to retrieve additional pages
* `parent` (string, optional) - Parent schema to list within, in the form
  `schemas/{catalog}.{schema}`. Required. Each `{...}` component is capped at
  255 characters individually
* `view` (string, optional) - Fields to return for each service. `FULL` includes destinations,
  inference-table details, and rate-limit principal names. `BASIC` omits
  destinations and inference-table details and omits principal names from
  rate limits. Defaults to `BASIC` when unset. Possible values are: `BASIC`, `FULL`
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.


## Attributes
This data source exports a single attribute, `model_services`. It is a list of resources, each with the following attributes:
* `comment` (string) - User-provided description
* `config` (ModelServiceConfig) - Destinations, routing, rate limits, and payload logging configuration.
  Required on Create. On Update, provide this field when `update_mask`
  contains `config` or one of its subpaths
* `create_time` (string) - Time the model service was created
* `created_by` (string) - Creator identity
* `effective_owner` (string) - Owner of the model service
* `etag` (string) - Optimistic concurrency token returned on every read. To make an Update or
  Delete conditional, pass the last-read value in that request's `etag`
  field. In REST responses, this value is a base64 string; URL-encode it when
  setting the `etag` query parameter
* `metastore_id` (string) - Metastore hosting the model service
* `name` (string) - Resource name of the model service.
  Format: `model-services/{catalog}.{schema}.{model_service}`.
  Each `{...}` component is capped at 255 characters individually.
  Server-derived on Create from `parent` +
  `model_service_id`; required and immutable on Update/Get/Delete
* `supported_api_types` (list of string) - API types supported across this service's destinations, such as
  `openai/v1/chat/completions`, `openai/v1/embeddings`, and
  `mlflow/v1/chat/completions`. Derived from the backing models and providers
  at read time
* `update_time` (string) - Time the model service was last modified
* `updated_by` (string) - Identity of the last updater

### InferenceTableConfig
* `is_deleted` (boolean) - Whether the referenced inference table has been deleted. The configuration
  remains visible so you can identify the broken dependency. Payload logging
  cannot continue until the table is restored or the configuration is updated
* `parent` (string) - Parent Unity Catalog schema where the inference table is created, in the
  form `schemas/{catalog}.{schema}`. Required when configuring an inference
  table. After the inference table is created, this field cannot be changed
* `table` (string) - Resolved UC table for payload logs.
  Format: `tables/{catalog}.{schema}.{table}`
* `table_name_prefix` (string) - Prefix used to form the inference table's registered name. AI Gateway
  appends `_payload`; for example, `table_name_prefix = "orders"` creates
  `orders_payload`. If unset, the prefix defaults to the service name. Read
  `table` from the response for the resulting resource name. After the
  inference table is created, this field cannot be changed

### ModelProviderServiceConfigModelTargetConfig
* `model` (string) - Provider-side model identifier, such as `gpt-5` or `claude-opus-4-7`.
  This identifies a model at the upstream provider; it is not a Unity
  Catalog model resource
* `native_api_types` (list of string) - Provider-native API types supported by this model, such as
  `openai/v1/chat/completions`. At least one value is required. AI Gateway
  uses these values to translate requests and responses. At most 64 entries
  of 256 characters each are allowed

### ModelServiceConfig
* `inference_table` (InferenceTableConfig) - Inference table configuration for payload logging
* `rate_limits` (list of RateLimit) - Rate limits applied to requests routed through this model service
* `routing` (ModelServiceConfigRoutingConfig) - Routing configuration: destinations and fallback

### ModelServiceConfigDestinationConfig
* `destination_type` (string) - Backing-model category. Provide the matching type-specific configuration
  and leave the other type-specific configurations unset. Possible values are: `DESTINATION_TYPE_EXTERNAL_FOUNDATION_MODEL`, `DESTINATION_TYPE_PAY_PER_TOKEN_FOUNDATION_MODEL`, `DESTINATION_TYPE_PROVISIONED_THROUGHPUT_FOUNDATION_MODEL`
* `external_model_config` (ModelServiceConfigExternalModelConfig) - Configuration for an external model reached through a model provider service
* `is_deleted` (boolean) - Whether the destination's backing model or model provider service has
  been deleted. The destination remains visible so you can identify the
  broken dependency. Requests cannot use this destination until the backing
  resource is restored or the destination is replaced
* `name` (string) - User-facing label for this destination, used in routing references
* `pay_per_token_config` (ModelServiceConfigPayPerTokenConfig) - Configuration for a pay-per-token Databricks foundation model
* `provisioned_throughput_config` (ModelServiceConfigProvisionedThroughputConfig) - Configuration for a provisioned-throughput Databricks foundation model
* `traffic_percentage` (integer) - Percentage of primary traffic sent to this destination, from 0 to 100.
  Required when there is more than one primary destination, in which case the
  primary percentages must sum to 100; a single primary destination receives
  all traffic. Fallback destinations are ordered and do not use this field

### ModelServiceConfigExternalModelConfig
* `model_provider_service` (string) - Resource name of the governed ModelProviderService that owns provider
  auth and provider-specific configuration. The referenced
  ModelProviderService also carries the provider type, so this message
  does not surface it directly.
  Format: `model-provider-services/{catalog}.{schema}.{model_provider_service}`.
  Each `{...}` component is capped at 255 characters individually
* `target` (ModelProviderServiceConfigModelTargetConfig) - Routing target for the destination: the provider-side model selected from
  the referenced ModelProviderService's `targets` catalog, plus the unified
  API types the platform should translate to/from at request time

### ModelServiceConfigFallbackConfig
* `destinations` (list of ModelServiceConfigDestinationConfig) - Fallback destinations, tried in the listed order. At most 5 are allowed

### ModelServiceConfigPayPerTokenConfig
* `model` (string) - Resource name of the Unity Catalog model.
  Format: `models/{catalog}.{schema}.{model}`

### ModelServiceConfigProvisionedThroughputConfig
* `model` (string) - UC model FQN of the model served by the backing endpoint (e.g.,
  `system.ai.databricks-claude-opus-4-6`). Resolved from Model Serving at
  Create/Update time
* `model_serving_endpoint` (string) - Name of the backing Model Serving endpoint serving the provisioned-
  throughput foundation model, in the form `serving-endpoints/{name}`. The
  same Unity Catalog model can be served on multiple Model Serving endpoints
  with different throughput, regions, or configurations. The caller selects
  the endpoint to which this destination routes. The endpoint must exist at
  create time

### ModelServiceConfigRoutingConfig
* `destinations` (list of ModelServiceConfigDestinationConfig) - Primary routing destinations. At most 10 are allowed. At least one is
  required on Create. On Update, provide this list when replacing the full
  `config` or updating `config.routing.destinations`; other granular routing
  updates do not require resending destinations. The intermediate
  `config.routing` mask path is not supported
* `fallback` (ModelServiceConfigFallbackConfig) - Fallback routing applied after a primary destination fails. Fallback
  destinations are tried in the listed order

### RateLimit
* `key` (string) - Scope of the rate limit. Depending on this value, the limit applies to a
  principal, the service as a whole, or each user by default. Possible values are: `RATE_LIMIT_KEY_SERVICE`, `RATE_LIMIT_KEY_SERVICE_PRINCIPAL`, `RATE_LIMIT_KEY_USER`, `RATE_LIMIT_KEY_USER_DEFAULT`, `RATE_LIMIT_KEY_USER_GROUP`
* `principal` (string) - Principal this limit applies to: user email, group name, or service
  principal application ID. Required when `key` applies to a user, group, or
  service principal; otherwise it must be unset
* `renewal_period` (string) - Renewal period. Possible values are: `RATE_LIMIT_RENEWAL_PERIOD_HOUR`, `RATE_LIMIT_RENEWAL_PERIOD_MINUTE`
* `requests` (integer) - Maximum requests allowed in one renewal period. Leave unset for no request
  limit. Set to `0` to deny all requests
* `tokens` (integer) - Maximum tokens allowed in one renewal period. Leave unset for no token
  limit. Set to `0` to deny all requests