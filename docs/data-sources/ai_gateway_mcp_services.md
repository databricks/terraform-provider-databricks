---
subcategory: "Unity Catalog"
---
# databricks_ai_gateway_mcp_services Data Source
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/aigateway)



## Example Usage


## Arguments
The following arguments are supported:
* `include_browse` (boolean, optional) - Whether to include MCP services for which the principal can only access
  selective metadata
* `page_size` (integer, optional) - Maximum number of MCP services to return. Defaults to 100 when unset or 0;
  the maximum is 100. Use `next_page_token` to retrieve additional pages
* `parent` (string, optional) - Resource name of the parent schema to list within, as
  `schemas/{catalog}.{schema}`. Each `{...}` component is capped at 255
  characters individually
* `view` (string, optional) - View selector controlling which fields are populated per row. Possible values are: `BASIC`, `FULL`
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.


## Attributes
This data source exports a single attribute, `mcp_services`. It is a list of resources, each with the following attributes:
* `browse_only` (boolean) - Whether the caller sees only metadata available through the BROWSE
  privilege
* `comment` (string) - User-provided description
* `config` (McpServiceConfig) - Operational configuration: connection, tool selectors, rate limit.
  Required on CreateMcpService; on
  UpdateMcpService it is required only when `config` (or a `config.*`
  subpath) appears in `update_mask`
* `create_time` (string) - When the MCP service was created
* `created_by` (string) - Creator identity
* `effective_owner` (string) - The resolved owner of the MCP service. Falls back to the caller's identity
  when `owner` is not explicitly set on creation
* `etag` (string) - Optimistic concurrency control token. Server-generated from the
  entity's state and returned on every read. To use it as an if-match
  precondition on a mutation, echo the last-read value back via the dedicated
  `etag` field on the Update / Delete request; the server rejects the mutation
  if the stored etag differs
* `metastore_id` (string) - Metastore hosting the MCP service
* `name` (string) - Resource name of the MCP service.
  Format: `mcp-services/{catalog}.{schema}.{mcp_service}`.
  Each `{...}` component is capped at 255 characters individually.
  Server-derived on Create from `parent` +
  `mcp_service_id`; required and immutable on Update/Get/Delete
* `owner` (string) - The owner of the MCP service. Write-only; read owner via effective_owner
* `update_time` (string) - When the MCP service was last modified
* `updated_by` (string) - Identity of the last updater

### McpServiceConfig
* `include_tool_selectors` (list of string) - Glob or exact-match patterns selecting which tools from the MCP server
  to expose. Prefix match for patterns with `*`, exact match otherwise.
  An empty list means all tools are included. Per-element max 256 chars
* `rate_limits` (list of RateLimit) - Per-principal rate limits applied to tool invocations routed through this
  MCP service. Repeated to support per-USER / USER_GROUP / SERVICE_PRINCIPAL
  / SERVICE / USER_DEFAULT scopes simultaneously, mirroring the
  `ModelServiceConfig.rate_limits` shape. Empty when no rate limit is
  configured
* `source_connection` (McpServiceConfigSourceConnection) - UC Connection referencing the MCP server

### McpServiceConfigSourceConnection
* `is_deleted` (boolean)
* `name` (string)

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