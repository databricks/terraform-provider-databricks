---
subcategory: "Unity Catalog"
---
# databricks_ai_gateway_mcp_service Data Source
[![GA](https://img.shields.io/badge/Release_Stage-GA-green)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/aigateway)

Retrieves a Unity Catalog MCP service by its full resource name.


## Example Usage
The following example retrieves the MCP service named `knowledge_tools` from the `main.default` schema:

```hcl
data "databricks_ai_gateway_mcp_service" "example" {
  name = "mcp-services/main.default.knowledge_tools"
}

output "mcp_service_config" {
  value = data.databricks_ai_gateway_mcp_service.example.config
}
```


## Arguments
The following arguments are supported:
* `name` (string, required) - Resource name of the MCP service.
  Format: `mcp-services/{catalog}.{schema}.{mcp_service}`.
  Each `{...}` component is capped at 255 characters individually.
  Server-derived on Create from `parent` +
  `mcp_service_id`; required and immutable on Update/Get/Delete
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
The following attributes are exported:
* `comment` (string) - User-provided description
* `config` (McpServiceConfig) - Connection, tool selectors, and rate limits. Required on Create. On Update,
  provide this field when `update_mask` contains `config` or one of its
  subpaths
* `create_time` (string) - Time the MCP service was created
* `created_by` (string) - Creator identity
* `effective_owner` (string) - Owner of the MCP service
* `etag` (string) - Optimistic concurrency token returned on every read. To make an Update or
  Delete conditional, pass the last-read value in that request's `etag`
  field. In REST responses, this value is a base64 string; URL-encode it when
  setting the `etag` query parameter
* `metastore_id` (string) - Metastore hosting the MCP service
* `name` (string) - Resource name of the MCP service.
  Format: `mcp-services/{catalog}.{schema}.{mcp_service}`.
  Each `{...}` component is capped at 255 characters individually.
  Server-derived on Create from `parent` +
  `mcp_service_id`; required and immutable on Update/Get/Delete
* `update_time` (string) - Time the MCP service was last modified
* `updated_by` (string) - Identity of the last updater

### McpServiceConfig
* `include_tool_selectors` (list of string) - Tool names or prefix patterns to expose from the MCP server. Use exact
  tool names or prefix patterns such as `read_*`. An empty list exposes all
  tools. At most 1,024 selectors are allowed, and each selector can contain
  at most 256 characters
* `rate_limits` (list of RateLimit) - Rate limits for tool invocations. Supported scopes are user, group, service
  principal, the service as a whole, and each user by default. Request and
  token limits are supported. Empty when no rate limit is configured
* `source_connection` (McpServiceConfigSourceConnection) - Unity Catalog connection referencing the MCP server. Required on Create

### McpServiceConfigSourceConnection
* `is_deleted` (boolean) - Whether the referenced connection has been deleted. The MCP service keeps
  the reference so callers can identify the broken dependency; tool
  invocation fails until the source connection is updated
* `name` (string) - Resource name of the Unity Catalog connection used to access the MCP
  server, in the form `connections/{catalog}.{schema}.{connection}`

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