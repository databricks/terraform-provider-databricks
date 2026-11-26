---
subcategory: "Lakewatch"
---
# databricks_lakewatch_builtin_parser_version Data Source
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/lakewatch)

Reads a built-in Lakewatch parser version, including its YAML spec. Use a numbered canonical name to select a fixed version. Built-in versions are read-only. To customize a built-in version, use its spec as the starting point for a custom parser.

This data source can only be used with a workspace-level provider.


## Example Usage
Replace the example name with a version returned by `databricks_lakewatch_builtin_parser_versions`:

```hcl
data "databricks_lakewatch_builtin_parser_version" "logs" {
  name = "builtin-parsers/22222222-2222-4222-8222-222222222222/versions/1"
}

output "builtin_parser_spec" {
  value = data.databricks_lakewatch_builtin_parser_version.logs.spec
}
```


## Arguments
The following arguments are supported:
* `name` (string, required) - Resource name. Format:
  `builtin-parsers/{parser_id}/versions/{version}`
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
The following attributes are exported:
* `description` (string) - Free-form description of this parser, authored by Databricks
* `name` (string) - Resource name. Format:
  `builtin-parsers/{parser_id}/versions/{version}`
* `parser_name` (string) - The parent built-in's published leaf name, copied from the container on read.
  Denormalized so a caller that needs identity alongside the spec reads the
  version alone rather than the container as well. Populated by
  `GetBuiltinParserVersion`; List omits it, as List omits `spec`
* `spec` (string) - The parser spec: the field definitions for this log source, as
  a YAML document. Read-only; built-in specs are authored by Databricks, so
  no client-facing size cap is enforced here
* `tags` (list of ParserTag) - The parent built-in's tags, copied from the container on read. Rendering reads
  these for connector-specific details, such as the source table a Lakeflow
  Connect parser reads
* `version` (integer) - Version number

### ParserTag
* `key` (string) - The tag key, e.g. `cloud_provider`. Required: a keyless tag cannot be
  filtered on and is meaningless
* `value` (string) - The tag value, e.g. `aws`