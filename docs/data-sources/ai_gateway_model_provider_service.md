---
subcategory: "Unity Catalog"
---
# databricks_ai_gateway_model_provider_service Data Source
[![GA](https://img.shields.io/badge/Release_Stage-GA-green)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/aigateway)

Retrieves a Unity Catalog model provider service by its full resource name. Secret values are not returned.


## Example Usage
The following example retrieves the model provider service named `custom_provider` from the `main.default` schema:

```hcl
data "databricks_ai_gateway_model_provider_service" "example" {
  name = "model-provider-services/main.default.custom_provider"
}

output "provider_type" {
  value = data.databricks_ai_gateway_model_provider_service.example.config.provider_type
}
```


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
* `comment` (string) - User-provided description
* `config` (ModelProviderServiceConfig) - Provider authentication, exposed models, request-forwarding controls, rate
  limits, and payload logging. Required on Create. On Update, it is required
  only when `config` or one of its subpaths appears in `update_mask`
* `create_time` (string) - Time the provider service was created
* `created_by` (string) - Creator identity
* `effective_owner` (string) - Owner of the model provider service
* `etag` (string) - Optimistic concurrency token returned on every read. To make an Update or
  Delete conditional, pass the last-read value in that request's `etag`
  field. In REST responses, this value is a base64 string; URL-encode it when
  setting the `etag` query parameter
* `metastore_id` (string) - Metastore hosting the provider service
* `name` (string) - Resource name of the provider service.
  Format: `model-provider-services/{catalog}.{schema}.{model_provider_service}`.
  Each `{...}` component is capped at 255 characters individually.
  Server-derived on Create from `parent` +
  `model_provider_service_id`; required and immutable on Update/Get/Delete
* `update_time` (string) - Time the provider service was last modified
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

### ModelProviderServiceConfig
* `allow_all_targets` (boolean) - When true, accepts any model exposed by the upstream provider; `targets`
  is not required and does not restrict routability. When false, only
  models listed in `targets` are routable. Defaults to false
* `amazon_bedrock` (ModelProviderServiceConfigAmazonBedrockProviderConfig)
* `anthropic` (ModelProviderServiceConfigAnthropicProviderConfig)
* `azure_openai` (ModelProviderServiceConfigAzureOpenAiProviderConfig)
* `custom` (ModelProviderServiceConfigCustomProviderConfig)
* `forward_headers` (boolean) - Whether to forward incoming HTTP headers to the upstream provider. Defaults
  to false and is configured for the entire provider service, not per request.
  Upstream authentication is configured separately in the provider-specific
  configuration
* `forward_query_parameters` (boolean) - Whether to forward incoming query parameters to the upstream provider.
  Defaults to false and is configured for the entire provider service, not
  per request
* `forward_unmanaged_paths` (boolean) - Whether to proxy paths that AI Gateway does not recognize as configured
  provider-native API types. Defaults to false. When true, these paths are
  forwarded unchanged to the upstream provider. When false, only
  recognized API paths are served. Enabling this broadens the upstream API
  surface exposed through the provider service
* `gemini_enterprise` (ModelProviderServiceConfigGeminiEnterpriseProviderConfig)
* `inference_table` (InferenceTableConfig) - Payload logging configuration for requests sent directly to this provider
  service. Requests routed through a model service are captured by that model
  service's inference table instead
* `microsoft_foundry` (ModelProviderServiceConfigMicrosoftFoundryProviderConfig)
* `openai` (ModelProviderServiceConfigOpenAiProviderConfig)
* `provider_type` (string) - External model provider. Required on Create and immutable thereafter. Set
  the matching provider-specific configuration, such as `openai`,
  `azure_openai`, or `amazon_bedrock`. Possible values are: `EXTERNAL_MODEL_PROVIDER_TYPE_AMAZON_BEDROCK`, `EXTERNAL_MODEL_PROVIDER_TYPE_ANTHROPIC`, `EXTERNAL_MODEL_PROVIDER_TYPE_AZURE_OPENAI`, `EXTERNAL_MODEL_PROVIDER_TYPE_CUSTOM`, `EXTERNAL_MODEL_PROVIDER_TYPE_GEMINI_ENTERPRISE`, `EXTERNAL_MODEL_PROVIDER_TYPE_MICROSOFT_FOUNDRY`, `EXTERNAL_MODEL_PROVIDER_TYPE_OPENAI`
* `rate_limits` (list of RateLimit) - Rate limits for requests sent directly to this provider service. Requests
  routed through a model service use that model service's rate limits instead
* `targets` (list of ModelProviderServiceConfigModelTargetConfig) - Models and provider-native API types exposed by this provider service. Each
  entry must include at least one `native_api_types` value. When
  `allow_all_targets` is false, at least one entry is required and model
  service destinations can reference only listed models. When
  `allow_all_targets` is true, any upstream model is routable; entries in
  this list provide API-type metadata without restricting other models

### ModelProviderServiceConfigAmazonBedrockProviderConfig
* `direct` (ModelProviderServiceConfigAmazonBedrockProviderDirectConfig) - Amazon Bedrock region and authentication configuration

### ModelProviderServiceConfigAmazonBedrockProviderDirectConfig
* `aws_access_key` (ModelProviderServiceConfigAwsAccessKey) - AWS access-key-pair authentication. Set `access_key_id` and
  `secret_access_key.plaintext`. Mutually exclusive with
  `service_credential`
* `region` (string) - AWS region where the Bedrock endpoint is hosted (e.g., `us-east-1`).
  Required on Create
* `service_credential` (ModelProviderServiceConfigServiceCredential) - Reference to a Unity Catalog service credential authorizing Bedrock
  requests. On Create, supply `service_credential.name` in the form
  `credentials/{name}`. Required on Create when using service-credential
  authentication; mutually exclusive with `aws_access_key`. The credential
  is referenced by name; its value is not carried here. Only
  supported on AWS-hosted workspaces

### ModelProviderServiceConfigAnthropicProviderConfig
* `direct` (ModelProviderServiceConfigAnthropicProviderDirectConfig) - Direct authentication with an API key supplied in
  `direct.api_key.plaintext`. Required unless `relayed` is set
* `relayed` (ModelProviderServiceConfigAnthropicProviderRelayedConfig) - Relayed authentication. Each inference request supplies the caller's
  OAuth token, which is forwarded to Anthropic. No Anthropic credential is
  stored. Mutually exclusive with `direct`

### ModelProviderServiceConfigAnthropicProviderDirectConfig
* `api_key` (ModelProviderServiceConfigProviderSecret) - Anthropic API key. Required when creating the service. Supply the value
  in `api_key.plaintext`

### ModelProviderServiceConfigAnthropicProviderRelayedConfig

### ModelProviderServiceConfigAwsAccessKey
* `access_key_id` (string) - AWS access key ID. Required on Create when using access-key auth. Treated as
  username-equivalent (not a secret value): round-trips on reads and is
  scrubbed from audit logs
* `secret_access_key` (ModelProviderServiceConfigProviderSecret) - AWS secret access key paired with `access_key_id`. Required when creating
  a service with access-key authentication. Supply the value in
  `secret_access_key.plaintext`

### ModelProviderServiceConfigAzureOpenAiProviderConfig
* `direct` (ModelProviderServiceConfigAzureOpenAiProviderDirectConfig) - Azure OpenAI endpoint and authentication configuration

### ModelProviderServiceConfigAzureOpenAiProviderDirectConfig
* `api_key` (ModelProviderServiceConfigProviderSecret) - Azure OpenAI API key. Supply the value in `api_key.plaintext`. Mutually
  exclusive with Entra ID and Unity Catalog service credential
  authentication
* `base_url` (string) - Full Azure OpenAI endpoint base URL, e.g.
  `https://myresource.openai.azure.com`. Required on Create
* `entra_service_principal` (ModelProviderServiceConfigEntraServicePrincipal) - Entra ID service-principal authentication. Set `tenant_id`, `client_id`,
  and `client_secret.plaintext`. Mutually exclusive with `api_key` and
  `service_credential`
* `service_credential` (ModelProviderServiceConfigServiceCredential) - Reference to a Unity Catalog service credential authorizing Azure OpenAI
  requests. On Create, supply `service_credential.name` in the form
  `credentials/{name}`. Required on Create when using service-credential
  authentication; mutually exclusive with `api_key` and
  `entra_service_principal`. The credential is referenced by name; its value
  is not carried here. Only supported on Azure-hosted workspaces

### ModelProviderServiceConfigCustomProviderConfig
* `direct` (ModelProviderServiceConfigCustomProviderDirectConfig) - Endpoint and authentication configuration for the custom provider

### ModelProviderServiceConfigCustomProviderDirectConfig
* `api_key` (ModelProviderServiceConfigProviderSecret) - Bearer token forwarded in the `Authorization` header. Supply the value
  in `api_key.plaintext`
* `base_url` (string) - Endpoint URL of the OpenAI-compatible service (e.g.,
  `https://api.example.com/v1`). Required on Create

### ModelProviderServiceConfigEntraServicePrincipal
* `client_id` (string) - Entra ID client (application) ID. Required on Create
* `client_secret` (ModelProviderServiceConfigProviderSecret) - Entra ID client secret. Supply the value in `client_secret.plaintext`
* `tenant_id` (string) - Entra ID (Azure AD) tenant ID. Required on Create

### ModelProviderServiceConfigGeminiEnterpriseProviderConfig
* `direct` (ModelProviderServiceConfigGeminiEnterpriseProviderDirectConfig) - Gemini Enterprise project, region, and authentication configuration

### ModelProviderServiceConfigGeminiEnterpriseProviderDirectConfig
* `api_key` (ModelProviderServiceConfigProviderSecret) - Google Gemini Enterprise API key. Required when creating the service.
  Supply the value in `api_key.plaintext`
* `project_id` (string) - GCP project ID hosting the Gemini Enterprise endpoint. Required on Create
* `region` (string) - GCP region of the Gemini Enterprise endpoint (e.g., `us-central1`).
  Required on Create

### ModelProviderServiceConfigMicrosoftFoundryProviderConfig
* `direct` (ModelProviderServiceConfigMicrosoftFoundryProviderDirectConfig) - Microsoft Foundry endpoint and authentication configuration

### ModelProviderServiceConfigMicrosoftFoundryProviderDirectConfig
* `api_key` (ModelProviderServiceConfigProviderSecret) - Microsoft Foundry API key. Supply the value in `api_key.plaintext`.
  Mutually exclusive with Entra ID and Unity Catalog service credential
  authentication
* `base_url` (string) - Microsoft Foundry endpoint URL. Required on Create
* `entra_service_principal` (ModelProviderServiceConfigEntraServicePrincipal) - Entra ID service-principal authentication. Set `tenant_id`, `client_id`,
  and `client_secret.plaintext`. Mutually exclusive with `api_key` and
  `service_credential`
* `service_credential` (ModelProviderServiceConfigServiceCredential) - Reference to a Unity Catalog service credential authorizing Microsoft
  Foundry requests. On Create, supply `service_credential.name` in the form
  `credentials/{name}`. Required on Create when using service-credential
  authentication; mutually exclusive with `api_key` and
  `entra_service_principal`. The credential is referenced by name; its value
  is not carried here. Only supported on Azure-hosted workspaces

### ModelProviderServiceConfigModelTargetConfig
* `model` (string) - Provider-side model identifier, such as `gpt-5` or `claude-opus-4-7`.
  This identifies a model at the upstream provider; it is not a Unity
  Catalog model resource
* `native_api_types` (list of string) - Provider-native API types supported by this model, such as
  `openai/v1/chat/completions`. At least one value is required. AI Gateway
  uses these values to translate requests and responses. At most 64 entries
  of 256 characters each are allowed

### ModelProviderServiceConfigOpenAiProviderConfig
* `direct` (ModelProviderServiceConfigOpenAiProviderDirectConfig) - OpenAI configuration with an API key supplied in the request

### ModelProviderServiceConfigOpenAiProviderDirectConfig
* `api_key` (ModelProviderServiceConfigProviderSecret) - OpenAI API key. Required when creating the service. Supply the value in
  `api_key.plaintext`
* `base_url` (string) - Optional custom base URL. Defaults to `https://api.openai.com/v1`. Use for
  OpenAI-API-compatible third-party endpoints or in-network proxies
* `organization` (string) - Optional OpenAI organization ID. When set, the platform forwards it as
  the `OpenAI-Organization` header

### ModelProviderServiceConfigProviderSecret
* `plaintext` (string) - Inline plaintext credential. INPUT_ONLY: the value never round-trips on
  reads. Get and List responses omit `plaintext`; the enclosing secret
  object remains present to indicate that a secret is configured

### ModelProviderServiceConfigServiceCredential
* `name` (string) - Resource name of the bound Unity Catalog service credential, in the form
  `credentials/{name}`. Supply this field when creating the service or
  rebinding its credential. On read, it reflects the credential's current
  name

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