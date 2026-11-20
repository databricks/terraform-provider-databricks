---
subcategory: "Lakewatch"
---
# databricks_lakewatch_builtin_parsers Data Source
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/lakewatch)



## Example Usage


## Arguments
The following arguments are supported:
* `filter` (string, optional) - AIP-160 filter expression. Built-ins are filterable by their `tags`; the
  exact tag-filter grammar is defined by the server
* `page_size` (integer, optional) - Maximum number of results to return in one page. This is a best-effort upper
  bound, not a guarantee: the server may return fewer items than requested (even
  zero) while more results still exist. If unset or 0, the server picks a default
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.


## Attributes
This data source exports a single attribute, `builtin_parsers`. It is a list of resources, each with the following attributes:
* `latest_version` (integer) - Newest published version number
* `name` (string) - Resource name. Format: `builtin-parsers/{parser_id}`
* `parser_id` (string) - The parser's UUID, authored into the built-in's published definition rather
  than minted per customer, so every workspace addresses the same built-in by
  the same id. It is the leaf of `name`
* `parser_name` (string) - The built-in's own leaf name, as published in content-marketplace. Stable and
  human-readable, which is what a UI shows; `name` addresses the resource by id
* `tags` (list of ParserTag) - Key/value tags describing this parser, e.g.
  `cloud_provider=aws` or `log_source=cloudtrail`. A parser may carry up to
  32. Filterable on list

### ParserTag
* `key` (string) - The tag key, e.g. `cloud_provider`. Required: a keyless tag cannot be
  filtered on and is meaningless
* `value` (string) - The tag value, e.g. `aws`