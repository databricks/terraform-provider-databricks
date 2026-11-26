---
subcategory: "Lakewatch"
---
# databricks_lakewatch_parser Resource
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/lakewatch)

Manages a custom Lakewatch parser. Creating a parser also creates version 1. The parser's `spec` and `description` supply the initial version only; later versions have their own `databricks_lakewatch_parser_version` configurations.

This resource can only be used with a workspace-level provider.

### Version lifecycle

The parser and each version are immutable. Changing the parser's configured creation fields requires replacing the parser, which gives it a new ID. To append a version while keeping the parser and its existing versions, add a `databricks_lakewatch_parser_version` configuration. Existing Lakewatch datasource references stay pinned to their chosen versions.

Deleting the parser deletes all of its remaining versions, including versions created outside Terraform. Lakewatch rejects deletion while any datasource references one of those versions.

### Importing an existing parser

The API does not return the parser's initial `spec` or `description`. When importing, configure the parser's actual catalog, schema, and parser name, retain a valid `spec`, and add `lifecycle.ignore_changes = [spec, description]`. Terraform still requires `spec` in the configuration. It uses this value if it needs to create the parser again; ignoring changes does not recover the original version 1 content.

If version 1 still exists, read its content with a `databricks_lakewatch_parser_version` data block. Importing the parser also works after version 1 has been deleted. The initial inputs do not represent the latest version or recreate a deleted version 1.

### Deleting a parser after deleting a Lakewatch datasource

Deleting a datasource can return success before Lakewatch finishes removing its parser references. A Terraform destroy can then stop with a reference-conflict error. Wait for Lakewatch to finish deleting the datasource, review the remaining plan, and rerun the destroy. The provider does not automatically retry this reference conflict. Other changes in the same apply or destroy may already have completed.


## Example Usage
Create a parser and its initial version:

```hcl
resource "databricks_lakewatch_parser" "logs" {
  catalog_name = "security_logs"
  schema_name  = "default"
  parser_name  = "application_logs"
  description  = "Initial parser version"
  spec         = <<-YAML
    fields:
      - name: source_ip
        type: STRING
        observables: [ip]
    time:
      coalesce:
        - path: ts
  YAML
}

output "parser_name" {
  value = databricks_lakewatch_parser.logs.name
}
```

### Import an existing parser

Replace the example name and catalog fields with the existing parser's values. Use a valid creation spec even if version 1 has been deleted.

```hcl
resource "databricks_lakewatch_parser" "imported" {
  catalog_name = "security_logs"
  schema_name  = "default"
  parser_name  = "existing_application_logs"
  spec         = <<-YAML
    fields:
      - name: source_ip
        type: STRING
        observables: [ip]
    time:
      coalesce:
        - path: ts
  YAML

  lifecycle {
    ignore_changes = [spec, description]
  }
}
```

Import by the parser's canonical name, then inspect `terraform plan`:

```sh
terraform import databricks_lakewatch_parser.imported parsers/11111111-1111-4111-8111-111111111111
terraform plan
```


## Arguments
The following arguments are supported:
* `catalog_name` (string, required) - Name of the parent UC catalog, set on create. Reads return the value as
  supplied; the current UC catalog name is reported in
  `effective_catalog_name`
* `parser_name` (string, required) - Client-supplied parser name, stored as supplied and fixed at create time.
  Must be a valid UC identifier. A custom parser is not currently a UC securable
  and this field has no rename path
* `schema_name` (string, required) - Name of the parent UC schema, set on create. Reads return the value as
  supplied; the current UC schema name is reported in `effective_schema_name`
* `spec` (string, required) - Spec of version 1, as a YAML document. Accepted only by `CreateParser`; read
  a version's spec with `GetParserVersion`
* `description` (string, optional) - Description of version 1. Accepted only by `CreateParser`; read a version's
  description with `GetParserVersion`
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
* `create_time` (string) - When the parser was created
* `created_by` (Principal) - The principal that created this parser. Each version records its own
  `created_by`
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

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "name"
  to = databricks_lakewatch_parser.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_lakewatch_parser.this "name"
```