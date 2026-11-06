---
subcategory: "Lakewatch"
---
# databricks_lakewatch_datasource_config Resource
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/lakewatch)



## Example Usage


## Arguments
The following arguments are supported:
* `parent` (string, required) - The resource name of the parent
* `default_compute_mode` (string, optional) - Default compute mode applied to datasource ingestion jobs when the
  individual datasource doesn't override it. v1
  `datasources.defaultComputeMode`. Possible values are: `DS_COMPUTE_PERFORMANCE_OPTIMIZED`, `DS_COMPUTE_STANDARD`
* `default_folder` (string, optional) - Schema name of the registered folder (in this catalog) that a new datasource
  is placed in when the request does not specify a schema. Must name a schema
  registered as a folder. Omit to clear. Reads return the value as supplied; the
  folder's current schema name is reported in `effective_default_folder`
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
* `effective_default_folder` (string) - The current schema name of the `default_folder` folder. Tracks UC schema
  renames. Unset when no default is configured, or when the folder's schema no
  longer exists or is not visible to the caller
* `etag` (string) - Opaque entity tag for optimistic concurrency. Returned on every read;
  pass it back as the Update request's `etag` to make the update
  conditional on the resource being unchanged
* `name` (string)
* `update_time` (string)
* `updated_by` (Principal)
* `version` (integer) - Monotonic version, bumped on each update. Informational; use `etag` for
  optimistic concurrency

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "name"
  to = databricks_lakewatch_datasource_config.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_lakewatch_datasource_config.this "name"
```