---
subcategory: "Lakewatch"
---
# databricks_lakewatch_parser Data Source
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/lakewatch)



## Example Usage


## Arguments
The following arguments are supported:
* `name` (string, required) - Resource name. Format: `parsers/{parser_id}`, where the leaf is the
  immutable server-assigned UUID. Create-time placement names are stored
  separately and do not affect the resource name
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
The following attributes are exported:
* `catalog_name` (string) - Name of the parent UC catalog, set on create. Reads return the value as
  supplied; the current UC catalog name is reported in
  `effective_catalog_name`
* `create_time` (string) - When the parser was created
* `created_by` (Principal) - The principal that created this parser. Each version records its own
  `created_by`
* `description` (string) - Description of version 1. Accepted only by `CreateParser`; read a version's
  description with `GetParserVersion`
* `effective_catalog_name` (string) - The current name of the parser's UC catalog. Tracks UC renames; when the
  current name cannot be looked up, the name as of the parser's last write
* `effective_schema_name` (string) - The current name of the parser's UC schema. Tracks UC renames; when the
  current name cannot be looked up, the name as of the parser's last write
* `latest_version` (integer) - Newest version of this parser. Read its spec by passing this to
  `GetParserVersion`, or by using the `latest` alias. Absent when every version
  has been deleted
* `name` (string) - Resource name. Format: `parsers/{parser_id}`, where the leaf is the
  immutable server-assigned UUID. Create-time placement names are stored
  separately and do not affect the resource name
* `owner` (Principal) - The principal that owns this parser. Server-set from the authenticated
  creator, never supplied by the client. Beta reads by UUID perform no
  authorization. Owner-gated write RPCs are outside this API's current
  implementation scope
* `parser_id` (string) - The parser's server-assigned UUID. It is not currently a UC securable ID,
  is immutable for the life of the resource, is the leaf of `name`, and is
  returned by CreateParser
* `parser_name` (string) - Client-supplied parser name, stored as supplied and fixed at create time.
  Must be a valid UC identifier. A custom parser is not currently a UC securable
  and this field has no rename path
* `schema_name` (string) - Name of the parent UC schema, set on create. Reads return the value as
  supplied; the current UC schema name is reported in `effective_schema_name`
* `spec` (string) - Spec of version 1, as a YAML document. Accepted only by `CreateParser`; read
  a version's spec with `GetParserVersion`

### Principal
* `display_name` (string) - Display name (user or SP name)
* `email` (string) - Email (regular users only)
* `id` (string) - User / service-principal ID
* `principal_type` (string) - Whether this principal is a regular user, service principal, or system principal. Possible values are: `REGULAR_USER`, `SERVICE_PRINCIPAL`, `SYSTEM_USER`