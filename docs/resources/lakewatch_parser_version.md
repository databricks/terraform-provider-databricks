---
subcategory: "Lakewatch"
---
# databricks_lakewatch_parser_version Resource
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/lakewatch)

Manages one immutable version of a custom Lakewatch parser. Creation appends a version to an existing parser. Lakewatch assigns its number; the Terraform configuration's label does not select a version number.

This resource can only be used with a workspace-level provider.

### Append or replace a version

Adding a new configuration creates a new version. Keep the earlier version's configuration if Terraform should continue managing it.

Changing `spec` or `description` replaces the tracked version. By default, Terraform deletes the tracked version before creating a new version with a higher number. Removing the configuration deletes the tracked version. Deleted numbers are never reused. Appending a version leaves the parser ID, earlier versions, and existing Lakewatch datasource references unchanged.

### Reference protection and empty parsers

Lakewatch rejects the deletion of any parser version referenced by a Lakewatch datasource, whether or not Terraform manages that datasource. Because replacement includes deleting the old version, Terraform cannot complete the replacement while a Lakewatch datasource still references that version. Remove those references by updating or deleting the affected datasources, then retry. Terraform does not roll back other operations that completed in the same apply.

An unreferenced version can be deleted even if it is version 1 or the last version. Deleting the last version leaves the parser empty. Appending to that parser continues the version sequence.

A datasource deletion can return success before its parser references are removed. If deleting the parser version fails with a reference-conflict error before Lakewatch finishes deleting the datasource, wait for cleanup to finish, review the remaining plan, and rerun the operation. The provider does not automatically retry this reference conflict.

### Importing a version

Import by the canonical numbered name, such as `parsers/{parser_id}/versions/2`. Import reads an existing version; it does not append one. Version 1 can also be imported.

Configure `parent` with the matching parser name and set `spec` and `description` to the existing version's values. The first plan after import can show an in-place change that records `parent` in Terraform state. Applying that change does not modify the version. A subsequent plan should show no changes.

The parser's initial inputs are creation values. Importing version 1 separately makes its deletion or replacement the responsibility of the version configuration; deleting it does not cause the parser configuration to recreate it.


## Example Usage
Append a version to an existing parser. Replace the parser name with the value returned when creating your parser:

```hcl
data "databricks_lakewatch_parser" "logs" {
  name = "parsers/11111111-1111-4111-8111-111111111111"
}

resource "databricks_lakewatch_parser_version" "next" {
  parent      = data.databricks_lakewatch_parser.logs.name
  description = "Add source IP extraction"
  spec        = <<-YAML
    fields:
      - name: source_ip
        type: STRING
        observables: [ip]
    time:
      coalesce:
        - path: ts
  YAML
}

output "parser_version_name" {
  value = databricks_lakewatch_parser_version.next.name
}
```

If Terraform also manages the parser, use `databricks_lakewatch_parser.logs.name` as `parent` to express that dependency.

### Import an existing version

Configure the version's actual parent, spec, and description before importing it into the configuration above:

```sh
terraform import databricks_lakewatch_parser_version.next parsers/11111111-1111-4111-8111-111111111111/versions/2
terraform plan
```


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