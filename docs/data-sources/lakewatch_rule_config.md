---
subcategory: "Lakewatch"
---
# databricks_lakewatch_rule_config Data Source
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/lakewatch)



## Example Usage


## Arguments
The following arguments are supported:
* `name` (string, required)
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
The following attributes are exported:
* `default_folder` (string) - Schema name of the registered folder (in this catalog) that a new rule is
  placed in when the request does not specify a schema. Must name a schema
  registered as a folder. Omit to clear. Reads return the value as supplied; the
  folder's current schema name is reported in `effective_default_folder`
* `default_schedule` (RuleSchedule) - Default schedule for rules that don't specify one
* `default_signal_collation_window` (string) - Default duration string for signal deduplication when a rule
  doesn't override it — compact day/hour/minute/second components, e.g.
  `5d`, `1h`. Reads return the value as supplied; there is no server default,
  so it is unset until set. An empty value clears it
* `detection_categories` (list of string) - Available detection categories a rule can be tagged with. Reads return the
  value as supplied. When empty, rules can be tagged with the default categories:
  `Static Signature`, `Fuzzy Signature`, `Behavioral`, `Statistical`, and
  `Learning (AI/ML)`. The categories in effect are reported in
  `effective_detection_categories`
* `effective_default_folder` (string) - The current schema name of the `default_folder` folder. Tracks UC schema
  renames. Unset when no default is configured, or when the folder's schema no
  longer exists or is not visible to the caller
* `effective_detection_categories` (list of string) - The detection categories in effect: `detection_categories` when set,
  otherwise the default categories
* `etag` (string) - Opaque entity tag for optimistic concurrency. Returned on every read;
  pass it back as the Update request's `etag` to make the update
  conditional on the resource being unchanged
* `name` (string)
* `update_time` (string)
* `updated_by` (Principal)
* `version` (integer) - Monotonic version, bumped on each update. Informational; use `etag` for
  optimistic concurrency

### Principal
* `display_name` (string) - Display name (user or SP name)
* `email` (string) - Email (regular users only)
* `id` (string) - User / service-principal ID
* `principal_type` (string) - Whether this principal is a regular user, service principal, or system principal. Possible values are: `REGULAR_USER`, `SERVICE_PRINCIPAL`, `SYSTEM_USER`

### RuleSchedule
* `at_least_every` (string) - Defines the minimum cadence as a duration string. It may contain components
  for days, hours, minutes, and seconds (e.g. `5d`, `5d3h`, `3d2h5m12s`,
  `12m16s`)
* `cron_expression` (string) - Defines the schedule using a Quartz cron expression
* `enabled` (boolean) - Sets whether the job is scheduled to run. Defaults to true