---
subcategory: "Lakewatch"
---
# databricks_lakewatch_parser_versions Data Source
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/lakewatch)

Lists the remaining versions of a custom Lakewatch parser, oldest first. Results contain version metadata without specs. Use `databricks_lakewatch_parser_version` to retrieve one version's content. Deleted versions are absent, so version numbers can have gaps.

This data source can only be used with a workspace-level provider.


## Example Usage
List the versions of an existing parser:

```hcl
data "databricks_lakewatch_parser_versions" "logs" {
  parent = "parsers/11111111-1111-4111-8111-111111111111"
}

output "parser_version_names" {
  value = [for version in data.databricks_lakewatch_parser_versions.logs.parser_versions : version.name]
}
```


## Arguments
The following arguments are supported:
* `parent` (string, required) - Resource name of the parser whose versions are listed. Format:
  `parsers/{parser_id}`
* `page_size` (integer, optional) - Maximum number of results to return in one page. This is a best-effort upper
  bound, not a guarantee: the server may return fewer items than requested (even
  zero) while more results still exist. If unset or 0, the server picks a default
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.


## Attributes
This data source exports a single attribute, `parser_versions`. It is a list of resources, each with the following attributes:
* `create_time` (string) - When this version was appended
* `created_by` (Principal) - The principal that appended this version
* `description` (string) - Free-form description of this version, set when the version is created
* `name` (string) - Resource name. Format:
  `parsers/{parser_id}/versions/{version}`
* `parser_name` (string) - The parent parser's name, copied from the container on read. Denormalized so a
  caller that needs identity alongside the spec reads the version alone rather
  than the container as well. Populated by `GetParserVersion`; List omits it, as
  List omits `spec`
* `spec` (string) - The field definitions for this log source, as a YAML document. Validated on
  write, so an invalid spec is never stored
* `version` (integer) - Version number, assigned by the server when the version is appended

### Principal
* `display_name` (string) - Display name (user or SP name)
* `email` (string) - Email (regular users only)
* `id` (string) - User / service-principal ID
* `principal_type` (string) - Whether this principal is a regular user, service principal, or system principal. Possible values are: `REGULAR_USER`, `SERVICE_PRINCIPAL`, `SYSTEM_USER`