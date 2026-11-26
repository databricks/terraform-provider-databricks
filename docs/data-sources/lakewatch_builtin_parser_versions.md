---
subcategory: "Lakewatch"
---
# databricks_lakewatch_builtin_parser_versions Data Source
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/lakewatch)

Lists published versions of a built-in Lakewatch parser, oldest first. Results contain version metadata without specs. Use `databricks_lakewatch_builtin_parser_version` to read one version's content.

This data source can only be used with a workspace-level provider.


## Example Usage
Replace the example parent with a name returned by `databricks_lakewatch_builtin_parsers`:

```hcl
data "databricks_lakewatch_builtin_parser_versions" "logs" {
  parent = "builtin-parsers/22222222-2222-4222-8222-222222222222"
}

output "builtin_parser_version_names" {
  value = [for version in data.databricks_lakewatch_builtin_parser_versions.logs.builtin_parser_versions : version.name]
}
```


## Arguments
The following arguments are supported:
* `parent` (string, required) - Resource name of the built-in parser whose versions are listed.
  Format: `builtin-parsers/{parser_id}`
* `page_size` (integer, optional) - Maximum number of results to return in one page. If unset or 0, at most 100 versions
  are returned. The maximum allowed value is 1000
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.


## Attributes
This data source exports a single attribute, `builtin_parser_versions`. It is a list of resources, each with the following attributes:
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