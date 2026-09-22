---
subcategory: "Unity Catalog"
---
# databricks_ai_gateway_mcp_service Resource
[![GA](https://img.shields.io/badge/Release_Stage-GA-green)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/aigateway)

Manages an MCP service in Unity Catalog. An MCP service governs access to an MCP server hosted through a Unity Catalog connection.

The Unity Catalog connection must exist before you create the MCP service. MCP services are contained in a Unity Catalog schema and governed by Unity Catalog permissions.


## Example Usage
The following example registers an MCP service backed by a Unity Catalog connection:

```hcl
resource "databricks_ai_gateway_mcp_service" "example" {
  parent         = "schemas/main.default"
  mcp_service_id = "knowledge_tools"
  comment        = "Provides governed access to knowledge tools"

  config = {
    source_connection = {
      name = "connections/main.default.mcp_connection"
    }
  }
}
```


## Arguments
The following arguments are supported:
* `mcp_service_id` (string, required) - Name for the MCP service, e.g. "my_mcp_service"
* `parent` (string, required) - Name of the parent schema.
  Format: `schemas/{catalog}.{schema}`.
  Each `{...}` component is capped at 255 characters individually
* `comment` (string, optional) - User-provided description
* `config` (McpServiceConfig, optional) - Connection, tool selectors, and rate limits. Required on Create. On Update,
  provide this field when `update_mask` contains `config` or one of its
  subpaths
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

### McpServiceConfig
* `include_tool_selectors` (list of string, optional) - Tool names or prefix patterns to expose from the MCP server. Use exact
  tool names or prefix patterns such as `read_*`. An empty list exposes all
  tools. At most 1,024 selectors are allowed, and each selector can contain
  at most 256 characters
* `rate_limits` (list of RateLimit, optional) - Rate limits for tool invocations. Supported scopes are user, group, service
  principal, the service as a whole, and each user by default. Request and
  token limits are supported. Empty when no rate limit is configured
* `source_connection` (McpServiceConfigSourceConnection, optional) - Unity Catalog connection referencing the MCP server. Required on Create

### McpServiceConfigSourceConnection
* `name` (string, required) - Resource name of the Unity Catalog connection used to access the MCP
  server, in the form `connections/{catalog}.{schema}.{connection}`

### RateLimit
* `key` (string, required) - Scope of the rate limit. Depending on this value, the limit applies to a
  principal, the service as a whole, or each user by default. Possible values are: `RATE_LIMIT_KEY_SERVICE`, `RATE_LIMIT_KEY_SERVICE_PRINCIPAL`, `RATE_LIMIT_KEY_USER`, `RATE_LIMIT_KEY_USER_DEFAULT`, `RATE_LIMIT_KEY_USER_GROUP`
* `renewal_period` (string, required) - Renewal period. Possible values are: `RATE_LIMIT_RENEWAL_PERIOD_HOUR`, `RATE_LIMIT_RENEWAL_PERIOD_MINUTE`
* `principal` (string, optional) - Principal this limit applies to: user email, group name, or service
  principal application ID. Required when `key` applies to a user, group, or
  service principal; otherwise it must be unset
* `requests` (integer, optional) - Maximum requests allowed in one renewal period. Leave unset for no request
  limit. Set to `0` to deny all requests
* `tokens` (integer, optional) - Maximum tokens allowed in one renewal period. Leave unset for no token
  limit. Set to `0` to deny all requests

## Attributes
In addition to the above arguments, the following attributes are exported:
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

### McpServiceConfigSourceConnection
* `is_deleted` (boolean) - Whether the referenced connection has been deleted. The MCP service keeps
  the reference so callers can identify the broken dependency; tool
  invocation fails until the source connection is updated

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "name"
  to = databricks_ai_gateway_mcp_service.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_ai_gateway_mcp_service.this "name"
```