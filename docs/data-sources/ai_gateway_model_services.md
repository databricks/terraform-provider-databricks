---
subcategory: "Unity Catalog"
---
# databricks_ai_gateway_model_services Data Source
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/aigateway)



## Example Usage


## Arguments
The following arguments are supported:
* `page_size` (integer, optional) - Maximum number of model services to return. Defaults to 100 when unset or 0;
  the maximum is 100. Use `page_token` to retrieve additional pages
* `parent` (string, optional) - Name of the parent schema to list within, as
  `schemas/{catalog}.{schema}`. Each `{...}` component is capped at 255
  characters individually
* `view` (string, optional) - View selector controlling which fields are populated per row. `FULL`
  returns the full representation of the service; `BASIC` returns a more
  compact version. Defaults to `BASIC` when unset. Possible values are: `BASIC`, `FULL`
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.


## Attributes
This data source exports a single attribute, `model_services`. It is a list of resources, each with the following attributes:
* `comment` (string) - User-provided description
* `config` (ModelServiceConfig) - Operational configuration: destinations, routing, rate limits, inference
  table. Required on CreateModelService; on UpdateModelService it is
  required only when `config` (or a `config.*` subpath) appears in
  `update_mask`
* `create_time` (string) - When the model service was created
* `created_by` (string) - Creator identity
* `effective_owner` (string) - The resolved owner of the ModelService. Falls back to the caller's identity
  when `owner` is not explicitly set on creation
* `etag` (string) - Optimistic concurrency control token. Server-generated from the
  entity's state and returned on every read. To use it as an if-match
  precondition on a mutation, echo the last-read value back via the dedicated
  `etag` field on the Update / Delete request; the server rejects the mutation
  if the stored etag differs
* `metastore_id` (string) - Metastore hosting the model service
* `name` (string) - Resource name of the model service.
  Format: `model-services/{catalog}.{schema}.{model_service}`.
  Each `{...}` component is capped at 255 characters individually.
  Server-derived on Create from `parent` +
  `model_service_id`; required and immutable on Update/Get/Delete
* `owner` (string) - The owner of the model service. Write-only; read owner via effective_owner
* `supported_api_types` (list of string) - Unified API types this endpoint supports (e.g. "chat", "embeddings",
  "completions"). Derived from the destinations' backing models / providers
  at read time
* `update_time` (string) - When the model service was last modified
* `updated_by` (string) - Identity of the last updater

### InferenceTableConfig
* `disabled` (boolean) - Indicates whether payload logging is disabled (opt-out). Unset means that
  payload logging is active (the on-by-default state coincides with the proto
  zero-value, so the server never fills this field for a client that leaves it
  unset). Set `disabled = true` to pause runtime logging while keeping the
  sub-message attached (preserving `parent` and `table_name_prefix` for a
  later flip back to active). `parent` remains required either way
* `is_deleted` (boolean) - True when the bound inference TABLE has been deleted but the parent
  service still references it. The dangling reference is surfaced (not
  silently dropped) so callers can see the broken dependency. AI Gateway
  payload logging fails closed in this state
* `parent` (string) - Parent UC schema where the inference table is created.
  Format: `schemas/{catalog}.{schema}`. Set at create time and immutable
  thereafter; changing it on an existing service is rejected
* `table` (string) - Resolved UC table for payload logs.
  Format: `tables/{catalog}.{schema}.{table}`
* `table_name_prefix` (string) - Prefix for the inference-table's UC-registered name. The actual leaf name UC
  stores is `<table_name_prefix>_payload`; the `_payload` suffix is appended
  automatically. To find the actual UC table after Create, read the `table`
  field on the response. Defaults to `<model_service_name>_payload` when unset.
  Set at create time and immutable thereafter; changing it on an existing
  service is rejected

### ModelProviderServiceConfigModelTargetConfig
* `model` (string) - Provider-side model identifier (e.g. "gpt-5", "claude-opus-4-7"). This is
  a string on the LLM provider's side, not a UC entity. The UC governance
  hook for external destinations is the ModelProviderService referenced by
  `ExternalModelConfig.model_provider_service`, not the model itself
* `native_api_types` (list of string) - Provider-native API types the model supports (e.g.
  "openai/v1/chat/completions"). Used by the platform for request/response
  translation from the unified API type. At most 64 entries of at most 256
  characters each; the list is persisted into the destination binding's
  bounded storage envelope

### ModelServiceConfig
* `inference_table` (InferenceTableConfig) - Inference table config for payload logging
* `rate_limits` (list of RateLimit) - Rate limits applied to requests routed through this model service
* `routing` (ModelServiceConfigRoutingConfig) - Routing configuration: destinations, routing strategy, and fallback

### ModelServiceConfigDestinationConfig
* `destination_type` (string) - Backing-model category. Determines which oneof variant is populated. Possible values are: `DESTINATION_TYPE_EXTERNAL_FOUNDATION_MODEL`, `DESTINATION_TYPE_PAY_PER_TOKEN_FOUNDATION_MODEL`, `DESTINATION_TYPE_PROVISIONED_THROUGHPUT_FOUNDATION_MODEL`
* `external_model_config` (ModelServiceConfigExternalModelConfig)
* `is_deleted` (boolean) - True when the destination's backing UC entity (MODEL for foundation-model
  destinations, MODEL_PROVIDER_SERVICE for external destinations) has been
  deleted but the destination row still references it. The dangling
  destination is surfaced (not silently dropped) so callers can see the
  broken routing. Inference traffic through this destination fails closed
  (BAD_REQUEST / FAILED_PRECONDITION)
* `name` (string) - User-facing label for this destination, used in routing references
* `pay_per_token_config` (ModelServiceConfigPayPerTokenConfig)
* `provisioned_throughput_config` (ModelServiceConfigProvisionedThroughputConfig)
* `traffic_percentage` (integer) - Share of traffic sent to this destination, 0-100. Optional on fallback
  destinations; see FallbackConfig

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
* `destinations` (list of ModelServiceConfigDestinationConfig) - Ordered list of fallback destinations. Traversal is in list order; the
  attempt count is the length of the list. At most 5 are allowed

### ModelServiceConfigPayPerTokenConfig
* `model` (string) - Resource name of the UC model.
  Format: `models/{catalog}.{schema}.{model}`

### ModelServiceConfigProvisionedThroughputConfig
* `model` (string) - UC model FQN of the model served by the backing endpoint (e.g.,
  `system.ai.databricks-claude-opus-4-6`). Resolved from Model Serving at
  Create/Update time
* `model_serving_endpoint` (string) - Name of the backing Model Serving endpoint serving the provisioned-
  throughput foundation model, as the AIP-122 typed resource name
  `serving-endpoints/{name}`. The same UC model can be served on multiple
  Model Serving endpoints (different throughput / region / config); the
  caller picks which one this destination routes to. The endpoint must
  exist at create time

### ModelServiceConfigRoutingConfig
* `destinations` (list of ModelServiceConfigDestinationConfig) - Primary routing destinations. At most 10 are allowed. At least one is
  required on CreateModelService; on UpdateModelService it is required only
  when `config.routing` (or a `config.routing.*` subpath) appears in
  `update_mask`
* `fallback` (ModelServiceConfigFallbackConfig) - Fallback routing config, applied after primary destinations fail
* `first_token_timeout` (string) - Timeout for the first token of a streaming response. If a destination does
  not return its first token within this duration, AI Gateway aborts the
  attempt and fails over to the next destination. Applies to streaming
  requests only. Leave unset for no first-token timeout
* `traffic_splitting` (ModelServiceConfigRoutingConfigTrafficSplitting) - Marker message selecting request-based traffic splitting. Traffic is
  distributed according to each destination's traffic_percentage value;
  no configuration lives on this message itself

### ModelServiceConfigRoutingConfigTrafficSplitting

### RateLimit
* `key` (string) - Scope key. Determines whether `principal` is required. Possible values are: `RATE_LIMIT_KEY_REQUEST_TAG`, `RATE_LIMIT_KEY_SERVICE`, `RATE_LIMIT_KEY_SERVICE_PRINCIPAL`, `RATE_LIMIT_KEY_USER`, `RATE_LIMIT_KEY_USER_DEFAULT`, `RATE_LIMIT_KEY_USER_GROUP`
* `principal` (string) - Principal this limit applies to: user email, group name, or service
  principal application ID. Required unless `key` is
  `RATE_LIMIT_KEY_SERVICE`, `RATE_LIMIT_KEY_USER_DEFAULT`, or
  `RATE_LIMIT_KEY_REQUEST_TAG` (which must not set a principal)
* `renewal_period` (string) - Renewal period. Possible values are: `RATE_LIMIT_RENEWAL_PERIOD_HOUR`, `RATE_LIMIT_RENEWAL_PERIOD_MINUTE`
* `request_tag_key` (string) - Request tag key this limit applies to. Required when `key` is
  `RATE_LIMIT_KEY_REQUEST_TAG`, forbidden otherwise
* `request_tag_value` (string) - Request tag value this limit applies to. Only valid when `key` is
  `RATE_LIMIT_KEY_REQUEST_TAG`. Leave unset to apply the limit to every
  value of `request_tag_key` (an any-value default); a set value is a
  specific override for that value
* `requests` (integer) - Max requests allowed within a renewal period. Leave unset for no request limit
* `tokens` (integer) - Max tokens allowed within a renewal period. Leave unset for no token limit