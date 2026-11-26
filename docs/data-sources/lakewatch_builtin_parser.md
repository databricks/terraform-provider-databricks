---
subcategory: "Lakewatch"
---
# databricks_lakewatch_builtin_parser Data Source
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/lakewatch)

Reads the identity and latest-version pointer of a built-in Lakewatch parser published by Databricks. Use `databricks_lakewatch_builtin_parser_version` to retrieve a version's YAML spec. Built-in parsers cannot be modified or deleted.

This data source can only be used with a workspace-level provider.


## Example Usage
Replace the example name with one returned by `databricks_lakewatch_builtin_parsers`:

```hcl
data "databricks_lakewatch_builtin_parser" "logs" {
  name = "builtin-parsers/22222222-2222-4222-8222-222222222222"
}

output "parser_display_name" {
  value = data.databricks_lakewatch_builtin_parser.logs.parser_name
}
```


## Arguments
The following arguments are supported:
* `name` (string, required) - Resource name. Format: `builtin-parsers/{parser_id}`
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
The following attributes are exported:
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