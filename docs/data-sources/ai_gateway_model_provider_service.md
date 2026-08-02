---
subcategory: "Unity Catalog"
---
# databricks_ai_gateway_model_provider_service Data Source
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/aigateway)



## Example Usage


## Arguments
The following arguments are supported:
* `name` (string, required) - Resource name of the provider service.
  Format: `model-provider-services/{catalog}.{schema}.{model_provider_service}`.
  Each `{...}` component is capped at 255 characters individually.
  Server-derived on Create from `parent` +
  `model_provider_service_id`; required and immutable on Update/Get/Delete
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
The following attributes are exported:
* `browse_only` (boolean) - Whether the caller sees only metadata available through the BROWSE
  privilege
* `comment` (string) - User-provided description
* `config` (ModelProviderServiceConfig) - Behavioral configuration: provider connection, model catalog, and
  passthrough policy. See `ModelProviderServiceConfig` for the per-field
  contract. Required on CreateModelProviderService; on Update it is required
  only when `config` (or a `config.*` subpath) appears in `update_mask`
* `create_time` (string) - When the provider service was created
* `created_by` (string) - Creator identity
* `effective_owner` (string) - The resolved owner of the model provider service. Falls back to the
  caller's identity when `owner` is not explicitly set on creation
* `etag` (string) - Optimistic concurrency control token. Server-generated from the
  entity's state and returned on every read. To use it as an if-match
  precondition on a mutation, echo the last-read value back via the dedicated
  `etag` field on the Update / Delete request; the server rejects the mutation
  if the stored etag differs
* `metastore_id` (string) - Metastore hosting the provider service
* `name` (string) - Resource name of the provider service.
  Format: `model-provider-services/{catalog}.{schema}.{model_provider_service}`.
  Each `{...}` component is capped at 255 characters individually.
  Server-derived on Create from `parent` +
  `model_provider_service_id`; required and immutable on Update/Get/Delete
* `owner` (string) - The owner of the model provider service. Write-only; read owner via
  effective_owner
* `update_time` (string) - When the provider service was last modified
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

### ModelProviderServiceConfig
* `allow_all_targets` (boolean) - When true, accepts any model exposed by the upstream provider; `targets`
  is not required and does not restrict routability. When false, only
  models listed in `targets` are routable
* `amazon_bedrock` (ModelProviderServiceConfigAmazonBedrockProviderConfig)
* `anthropic` (ModelProviderServiceConfigAnthropicProviderConfig)
* `azure_openai` (ModelProviderServiceConfigAzureOpenAiProviderConfig)
* `custom` (ModelProviderServiceConfigCustomProviderConfig)
* `forward_headers` (boolean) - Whether to forward incoming request headers to the upstream provider.
  Applies to managed (multi-model) requests as well as passthrough requests
  served by this provider service. Governance-level decision by the provider
  service owner; not selectable per inference call
* `forward_query_parameters` (boolean) - Whether to forward incoming request query parameters to the upstream
  provider. Same trust-boundary semantics as `forward_headers`
* `forward_unmanaged_paths` (boolean) - Whether to forward request paths that fall outside this service's managed
  API set to the upstream provider as opaque passthrough. When true,
  requests addressed to subpaths not recognized by the managed API surface
  are proxied to the upstream provider over the same provider connection.
  When false, only managed-API paths are served. Governance-level decision
  by the provider service owner; expanding this expands the trust boundary
  that the ModelProviderService exposes
* `gemini_enterprise` (ModelProviderServiceConfigGeminiEnterpriseProviderConfig)
* `inference_table` (InferenceTableConfig) - Inference table configuration for payload logging when this provider
  service is invoked directly. When it is invoked through a model service,
  the model service's own inference table captures the invocation instead.
  Mirrors `ModelServiceConfig.inference_table` /
  `AgentServiceConfig.inference_table`
* `microsoft_foundry` (ModelProviderServiceConfigMicrosoftFoundryProviderConfig)
* `openai` (ModelProviderServiceConfigOpenAiProviderConfig)
* `provider_type` (string) - Provider type discriminator. Required at create time; immutable after.
  Determines which variant of the `provider` oneof must be set. May not be
  changed via Update; attempts to include `config.provider_type` in
  `UpdateModelProviderServiceRequest.update_mask` are rejected.
  
  Required on CreateModelProviderService and immutable thereafter. Possible values are: `EXTERNAL_MODEL_PROVIDER_TYPE_AMAZON_BEDROCK`, `EXTERNAL_MODEL_PROVIDER_TYPE_ANTHROPIC`, `EXTERNAL_MODEL_PROVIDER_TYPE_AZURE_OPENAI`, `EXTERNAL_MODEL_PROVIDER_TYPE_CUSTOM`, `EXTERNAL_MODEL_PROVIDER_TYPE_GEMINI_ENTERPRISE`, `EXTERNAL_MODEL_PROVIDER_TYPE_MICROSOFT_FOUNDRY`, `EXTERNAL_MODEL_PROVIDER_TYPE_OPENAI`
* `rate_limits` (list of RateLimit) - Rate limits applied when this provider service is invoked directly. When
  it is invoked through a model service, the model service's own
  `rate_limits` apply instead. Mirrors `ModelServiceConfig.rate_limits` /
  `McpServiceConfig.rate_limits`
* `targets` (list of ModelProviderServiceConfigModelTargetConfig) - Routing targets this provider service exposes (provider-side model
  identifier + unified API types per entry). Required (>=1) when
  `allow_all_targets = false`; optional and additive when
  `allow_all_targets = true`. References from `ExternalModelConfig.target`
  must match an entry here unless `allow_all_targets = true`

### ModelProviderServiceConfigAmazonBedrockProviderConfig
* `direct` (ModelProviderServiceConfigAmazonBedrockProviderDirectConfig)

### ModelProviderServiceConfigAmazonBedrockProviderDirectConfig
* `aws_access_key` (ModelProviderServiceConfigAwsAccessKey) - AWS access-key-pair auth. Mutually exclusive with `service_credential`.
  Supersedes the flat `aws_access_key_id` / `aws_secret_access_key` fields
* `aws_access_key_id` (string, deprecated) - Deprecated flat AWS access key ID. Superseded by
  `aws_access_key.access_key_id`. Kept for one migration cycle; the handler
  mirrors it to/from `aws_access_key`. Treated as username-equivalent (not a
  secret value): round-trips on reads and is scrubbed from audit logs
* `aws_secret_access_key` (ModelProviderServiceConfigProviderSecret, deprecated) - Deprecated flat AWS secret access key. Superseded by
  `aws_access_key.secret_access_key`. Kept for one migration cycle; the
  handler mirrors it to/from `aws_access_key`. Supplied as inline plaintext
  via `ProviderSecret.plaintext`
* `region` (string) - AWS region where the Bedrock endpoint is hosted (e.g., `us-east-1`).
  Required on Create
* `service_credential` (ModelProviderServiceConfigServiceCredential) - Reference to a UC service credential authorizing Bedrock requests. On
  Create the caller supplies `service_credential.name` in the AIP-122
  resource-name form `credentials/{name}`. Required on Create when using
  UC-service-credential auth; mutually exclusive with `aws_access_key`. The
  credential is referenced by name; its value is not carried here. On read the
  resolved `id` and `is_deleted` are also populated. Only supported on AWS-hosted
  workspaces; Create requests from other clouds are rejected with
  INVALID_PARAMETER_VALUE

### ModelProviderServiceConfigAnthropicProviderConfig
* `direct` (ModelProviderServiceConfigAnthropicProviderDirectConfig) - Direct (inline-credentials) form: caller supplies the API key in the
  request body. Required on Create unless `relayed` is set
* `relayed` (ModelProviderServiceConfigAnthropicProviderRelayedConfig) - Relayed (credential-less) form: no Anthropic credential is stored. Each
  inference request instead carries the caller's own OAuth token, which the
  platform forwards to Anthropic on outbound requests. Mutually exclusive
  with `direct`; no `api_key` is required or persisted

### ModelProviderServiceConfigAnthropicProviderDirectConfig
* `api_key` (ModelProviderServiceConfigProviderSecret) - Anthropic API key. Required on Create. Sent as the `x-api-key` header on
  outbound requests. Supplied as inline plaintext via
  `ProviderSecret.plaintext`

### ModelProviderServiceConfigAnthropicProviderRelayedConfig
* `plan_type` (string) - Which Anthropic subscription tier the relayed token belongs to. Optional;
  when unset the MPS gets the full governance surface (see TEAM_ENTERPRISE).
  Immutable after Create, so the tier cannot be flipped in place. Possible values are: `ANTHROPIC_RELAYED_PLAN_TYPE_MAX`, `ANTHROPIC_RELAYED_PLAN_TYPE_TEAM_ENTERPRISE`

### ModelProviderServiceConfigAwsAccessKey
* `access_key_id` (string) - AWS access key ID. Required on Create when using access-key auth. Treated as
  username-equivalent (not a secret value): round-trips on reads and is
  scrubbed from audit logs
* `secret_access_key` (ModelProviderServiceConfigProviderSecret) - AWS secret access key paired with `access_key_id`. Required on Create when
  using access-key auth. Supplied as inline plaintext via
  `ProviderSecret.plaintext`

### ModelProviderServiceConfigAzureOpenAiProviderConfig
* `direct` (ModelProviderServiceConfigAzureOpenAiProviderDirectConfig)

### ModelProviderServiceConfigAzureOpenAiProviderDirectConfig
* `api_key` (ModelProviderServiceConfigProviderSecret) - Azure OpenAI API key. Mutually exclusive with the Entra and
  service-credential modes. Supplied as inline plaintext via
  `ProviderSecret.plaintext`
* `base_url` (string) - Full Azure OpenAI endpoint base URL, e.g.
  `https://myresource.openai.azure.com`. Required on Create
* `client_id` (string, deprecated) - Deprecated flat Entra client ID. Superseded by
  `entra_service_principal.client_id`. Kept for one migration cycle; the
  handler mirrors it to/from `entra_service_principal`
* `client_secret` (ModelProviderServiceConfigProviderSecret, deprecated) - Deprecated flat Entra client secret. Superseded by
  `entra_service_principal.client_secret`. Kept for one migration cycle; the
  handler mirrors it to/from `entra_service_principal`. Supplied as inline
  plaintext via `ProviderSecret.plaintext`
* `entra_service_principal` (ModelProviderServiceConfigEntraServicePrincipal) - Entra ID (service principal) auth. Mutually exclusive with `api_key` and
  `service_credential`. Supersedes the flat `tenant_id` / `client_id` /
  `client_secret` fields
* `service_credential` (ModelProviderServiceConfigServiceCredential) - Reference to a UC service credential authorizing Azure OpenAI requests. On
  Create the caller supplies `service_credential.name` in the AIP-122
  resource-name form `credentials/{name}`. Required on Create when using
  UC-service-credential auth; mutually exclusive with `api_key` and
  `entra_service_principal`. The credential is
  referenced by name; its value is not carried here. On read the resolved `id`
  and `is_deleted` are also populated. Only supported on Azure-hosted
  workspaces; Create requests from other clouds are rejected with
  INVALID_PARAMETER_VALUE
* `tenant_id` (string, deprecated) - Deprecated flat Entra tenant ID. Superseded by
  `entra_service_principal.tenant_id`. Kept for one migration cycle; the
  handler mirrors it to/from `entra_service_principal`

### ModelProviderServiceConfigCustomProviderConfig
* `direct` (ModelProviderServiceConfigCustomProviderDirectConfig)

### ModelProviderServiceConfigCustomProviderDirectConfig
* `api_key` (ModelProviderServiceConfigProviderSecret) - Bearer token forwarded as the `Authorization: Bearer ...` header on
  outbound requests. Supplied as inline plaintext via
  `ProviderSecret.plaintext`. Set this for bearer-token auth
* `base_url` (string) - Endpoint URL of the OpenAI-compatible service (e.g.,
  `https://api.example.com/v1`). Required on Create

### ModelProviderServiceConfigEntraServicePrincipal
* `client_id` (string) - Entra ID client (application) ID. Required on Create
* `client_secret` (ModelProviderServiceConfigProviderSecret) - Entra ID client secret. Supplied as inline plaintext via
  `ProviderSecret.plaintext`
* `tenant_id` (string) - Entra ID (Azure AD) tenant ID. Required on Create

### ModelProviderServiceConfigGeminiEnterpriseProviderConfig
* `direct` (ModelProviderServiceConfigGeminiEnterpriseProviderDirectConfig)

### ModelProviderServiceConfigGeminiEnterpriseProviderDirectConfig
* `api_key` (ModelProviderServiceConfigProviderSecret) - Google Gemini Enterprise API key. Required on Create. Supplied as inline
  plaintext via `ProviderSecret.plaintext`
* `project_id` (string) - GCP project ID hosting the Gemini Enterprise endpoint. Required on Create
* `region` (string) - GCP region of the Gemini Enterprise endpoint (e.g., `us-central1`).
  Required on Create

### ModelProviderServiceConfigMicrosoftFoundryProviderConfig
* `direct` (ModelProviderServiceConfigMicrosoftFoundryProviderDirectConfig)

### ModelProviderServiceConfigMicrosoftFoundryProviderDirectConfig
* `api_key` (ModelProviderServiceConfigProviderSecret) - Microsoft AI Foundry API key. Mutually exclusive with the Entra and
  service-credential modes. Supplied as inline plaintext via
  `ProviderSecret.plaintext`
* `base_url` (string) - Microsoft AI Foundry endpoint URL. Required on Create
* `client_id` (string, deprecated) - Deprecated flat Entra client ID. Superseded by
  `entra_service_principal.client_id`. Kept for one migration cycle; the
  handler mirrors it to/from `entra_service_principal`
* `client_secret` (ModelProviderServiceConfigProviderSecret, deprecated) - Deprecated flat Entra client secret. Superseded by
  `entra_service_principal.client_secret`. Kept for one migration cycle; the
  handler mirrors it to/from `entra_service_principal`. Supplied as inline
  plaintext via `ProviderSecret.plaintext`
* `entra_service_principal` (ModelProviderServiceConfigEntraServicePrincipal) - Entra ID (service principal) auth. Mutually exclusive with `api_key` and
  `service_credential`. Supersedes the flat `tenant_id` / `client_id` /
  `client_secret` fields
* `service_credential` (ModelProviderServiceConfigServiceCredential) - Reference to a UC service credential authorizing Microsoft Foundry requests.
  On Create the caller supplies `service_credential.name` in the AIP-122
  resource-name form `credentials/{name}`. Required on Create when using
  UC-service-credential auth; mutually exclusive with `api_key` and
  `entra_service_principal`. The credential is
  referenced by name; its value is not carried here. On read the resolved `id`
  and `is_deleted` are also populated. Only supported on Azure-hosted
  workspaces; Create requests from other clouds are rejected with
  INVALID_PARAMETER_VALUE
* `tenant_id` (string, deprecated) - Deprecated flat Entra tenant ID. Superseded by
  `entra_service_principal.tenant_id`. Kept for one migration cycle; the
  handler mirrors it to/from `entra_service_principal`

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

### ModelProviderServiceConfigOpenAiProviderConfig
* `direct` (ModelProviderServiceConfigOpenAiProviderDirectConfig)

### ModelProviderServiceConfigOpenAiProviderDirectConfig
* `api_key` (ModelProviderServiceConfigProviderSecret) - OpenAI API key. Required on Create. Supplied as inline plaintext via
  `ProviderSecret.plaintext`
* `base_url` (string) - Optional custom base URL. Defaults to `https://api.openai.com/v1`. Use for
  OpenAI-API-compatible third-party endpoints or in-network proxies
* `organization` (string) - Optional OpenAI organization ID. When set, the platform forwards it as
  the `OpenAI-Organization` header

### ModelProviderServiceConfigProviderSecret
* `plaintext` (string) - Inline plaintext credential. INPUT_ONLY: the value never round-trips on
  reads. Get and List responses omit `plaintext`; the field's presence in
  the read shape only indicates that a secret is configured

### ModelProviderServiceConfigServiceCredential
* `name` (string) - Resource name of the bound UC service credential, in the AIP-122 form
  `credentials/{name}` (a metastore-level single-part credential name). On
  create the caller supplies the name here. On read it reflects the
  credential's current name at read time

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