---
subcategory: "Lakewatch"
---
# databricks_lakewatch_parser_version Data Source
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/lakewatch)

Reads a custom Lakewatch parser version, including its YAML spec and description. Use a numbered canonical name to select a fixed version. The `latest` alias selects the highest version that still exists and returns an error if the parser is empty.

This data source can only be used with a workspace-level provider.


## Example Usage
Read a fixed version:

```hcl
data "databricks_lakewatch_parser_version" "logs" {
  name = "parsers/11111111-1111-4111-8111-111111111111/versions/2"
}

output "parser_spec" {
  value = data.databricks_lakewatch_parser_version.logs.spec
}
```


## Arguments
The following arguments are supported:
* `name` (string, required) - Resource name. Format:
  `parsers/{parser_id}/versions/{version}`
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
The following attributes are exported:
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