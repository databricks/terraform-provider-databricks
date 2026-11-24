---
subcategory: "Lakewatch"
---
# databricks_lakewatch_parser_version Resource
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/lakewatch)



## Example Usage


## Arguments
The following arguments are supported:
* `parent` (string, required) - Resource name of the parser to append a version to. Format:
  `parsers/{parser_id}`
* `spec` (string, required) - The field definitions for this log source, as a YAML document. Validated on
  write, so an invalid spec is never stored
* `description` (string, optional) - Free-form description of this version, set when the version is created
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

### Principal
* `display_name` (string, optional) - Display name (user or SP name)
* `email` (string, optional) - Email (regular users only)
* `id` (string, optional) - User / service-principal ID
* `principal_type` (string, optional) - Whether this principal is a regular user, service principal, or system principal. Possible values are: `REGULAR_USER`, `SERVICE_PRINCIPAL`, `SYSTEM_USER`

## Attributes
In addition to the above arguments, the following attributes are exported:
* `create_time` (string) - When this version was appended
* `created_by` (Principal) - The principal that appended this version
* `name` (string) - Resource name. Format:
  `parsers/{parser_id}/versions/{version}`
* `parser_name` (string) - The parent parser's name, copied from the container on read. Denormalized so a
  caller that needs identity alongside the spec reads the version alone rather
  than the container as well. Populated by `GetParserVersion`; List omits it, as
  List omits `spec`
* `version` (integer) - Version number, assigned by the server when the version is appended

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "name"
  to = databricks_lakewatch_parser_version.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_lakewatch_parser_version.this "name"
```