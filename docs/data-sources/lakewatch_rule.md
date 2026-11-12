---
subcategory: "Lakewatch"
---
# databricks_lakewatch_rule Data Source
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/lakewatch)



## Example Usage


## Arguments
The following arguments are supported:
* `name` (string, required) - Resource name, in the form `rules/{rule_id}`. Assigned by the server on create; required
  on update
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
The following attributes are exported:
* `annotations` (object) - Free-form client-managed key/value annotations
* `batch` (RuleInputBatch) - Batch SQL rule input
* `catalog_name` (string) - Name of the Unity Catalog catalog that contains the rule. Required. Cannot be changed after
  creation. Reads return the value as supplied; the catalog's current name is reported in
  `effective_catalog_name`
* `category` (string) - Free-form rule category
* `compute_workspace_id` (string) - ID of the workspace whose compute runs this rule's jobs and owns its
  associated resources, such as the SQL warehouse a scheduled rule runs on.
  Server-assigned; a rule's warehouse can only be changed from this workspace
* `content_resource_id` (string) - Immutable UUIDv4 identifying the authored content resource installed as this rule
* `create_time` (string) - Time the rule was created
* `created_by` (Principal) - Principal that created the rule
* `deduplication` (string) - Signal deduplication mode. Possible values are: `DEDUPLICATION_DEFAULT`, `DEDUPLICATION_DISABLED`, `DEDUPLICATION_ENABLED`
* `deduplication_time_window` (string) - Time window for deduplication, such as `1h`, `24h`, `7d`, or `automatic`
* `description` (string) - Free-form description
* `effective_catalog_name` (string) - Current name of the Unity Catalog catalog that contains the rule. Reflects catalog renames,
  which do not change `catalog_name`. Unset until resolved
* `effective_rule_name` (string) - Current UC leaf name of the rule (server-owned). Tracks UC renames so callers
  can observe the live name without having to update the client-owned
  `rule_name`; unset until resolved
* `effective_run_as_principal` (string) - The identity this rule's job runs as: the resolved `run_as_principal`, or the
  requester's identity when it was omitted on create or cleared on update. The value is a user
  email or service-principal application ID, falling back to the numeric principal ID if that
  identity can no longer be resolved. Populated on single-rule reads (get/create/update); list
  responses omit it
* `effective_schema_name` (string) - Current name of the Unity Catalog schema that contains the rule. Reflects schema renames,
  which do not change `schema_name`. Unset until resolved
* `etag` (string) - Opaque entity tag for optimistic concurrency. Content-version ETags change on
  user edits and remain stable across runtime status updates
* `fidelity` (string) - Signal quality for this rule. Possible values are: `FIDELITY_DEVELOPMENT`, `FIDELITY_HIGH`, `FIDELITY_INVESTIGATIVE`, `FIDELITY_MEDIUM`
* `mitre` (list of RuleMitreMapping) - Up to 64 MITRE ATT&CK mappings, producing at most 192 flattened alert tags
* `name` (string) - Resource name, in the form `rules/{rule_id}`. Assigned by the server on create; required
  on update
* `objective` (string) - Long-form description of what this rule detects
* `original_content_pack` (ContentPackReference) - Reference to the content pack this rule was originally installed from
* `output` (RuleOutput) - What to record in each signal
* `owner` (Principal) - Unity Catalog owner of the rule. Assigned by the server. Populated by `GetRule`
* `references` (list of string) - Ordered reference URLs carried from the authored rule metadata. A maximum of three references
  of up to 512 characters each is supported. Requests exceeding these limits are rejected
* `rule_id` (string) - Unity Catalog securable UUID of the rule, and the last segment of `name`. Never changes,
  including when the rule's catalog or schema is renamed
* `rule_name` (string) - Name of the rule within its Unity Catalog schema. Required on create; can be updated
  within the same catalog and schema. Get synchronizes this field with the current Unity Catalog
  name, also reported in `effective_rule_name`
* `run_as_principal` (string) - The principal whose identity this rule's job runs as, given as a
  user email or service-principal application ID. When supplied on
  create/update, Lakewatch resolves the principal, records it on the rule, and
  creates the delegation to run as that principal — the requester must be
  authorized to delegate to it, or the request is rejected. When omitted on
  create, the requester's own identity is used. An update whose mask does not
  include `run_as_principal` keeps the current identity, and the requester must
  still be authorized to delegate to it; an editor who is not must supply a
  run_as_principal they can delegate to (for example, their own identity), or
  clear it, either of which reassigns the rule. An update whose mask includes
  `run_as_principal` but supplies no value clears it: the rule then runs as the
  requester, exactly as when it is omitted on create. Reads return the value as
  supplied; the identity the rule runs as is reported in
  `effective_run_as_principal`
* `runbook` (string) - Long-form operational guidance carried from the authored rule metadata. Supports up to 4096
  characters. Requests exceeding this limit are rejected. Supports Markdown; analysts and the
  investigation agent follow it when triaging this rule's alerts
* `schedule` (RuleSchedule) - When the rule executes
* `schema_name` (string) - Name of the Unity Catalog schema that contains the rule. Required. Cannot be changed after
  creation. Reads return the value as supplied; the schema's current name is reported in
  `effective_schema_name`
* `severity` (string) - Free-form severity label, such as `Informational`, `Medium`, `High`, or
  `Critical`
* `status` (ResourceStatus) - Read-only lifecycle status
* `suppress_until_time` (string) - If set, this rule's signals do not open new alerts until the given time
* `timerange` (RuleTimeRange) - Time-range configuration for batch rules. If unset, the execution time range is extended by
  a fixed `15m` lookback
* `update_time` (string) - Time the rule was last updated
* `updated_by` (Principal) - Principal that last updated the rule
* `version` (integer) - Monotonic version, bumped on each update. Informational; use `etag` for
  optimistic concurrency

### ContentPackReference
* `uuid` (string) - UUID of the content pack
* `version` (string) - Pack version the resource was installed from

### Principal
* `display_name` (string) - Display name (user or SP name)
* `email` (string) - Email (regular users only)
* `id` (string) - User / service-principal ID
* `principal_type` (string) - Whether this principal is a regular user, service principal, or system principal. Possible values are: `REGULAR_USER`, `SERVICE_PRINCIPAL`, `SYSTEM_USER`

### ResourceStatus
* `create_time` (string)
* `job_id` (string) - Databricks Jobs API job ID. Unset for a pure ingestion pipeline
* `last_complete_time` (string) - When the most recent run completed. Pairs with `run_state = SUCCEEDED`
* `last_error_message` (string) - Most recent error message, if any. Pairs with `run_state = FAILED`
* `last_start_time` (string) - When the most recent run started. Pairs with `run_state = RUNNING`
* `last_status_change_time` (string)
* `lifecycle_state` (string) - Scheduling lifecycle of the resource's underlying job/pipeline: whether it
  is unscheduled, scheduled, stopped, or deleted. Distinct from `run_state`,
  which reports the outcome of the most recent run. Possible values are: `DELETED`, `SCHEDULED`, `STOPPED`, `UNSCHEDULED`
* `pipeline_id` (string) - DLT ingestion-pipeline ID. Set only when this status represents the
  datasource's ingestion pipeline (unset for job statuses)
* `run_state` (string) - Possible values are: `CANCELED`, `FAILED`, `RUNNING`, `SUCCEEDED`

### RuleInputBatch
* `custom` (RuleInputCustomNotebook) - Custom-notebook user function. The output is used as input data for the
  rule
* `sql` (RuleInputBatchSql) - SQL query evaluated as the rule input
* `warehouse_id` (string) - SQL warehouse used to run this batch rule

### RuleInputBatchSql
* `default_catalog_name` (string) - Default catalog used to resolve unqualified identifiers in `query`. If
  unset, the rule's `catalog_name` is used
* `default_schema_name` (string) - Default schema used to resolve unqualified identifiers in `query`. If
  unset, the rule's `schema_name` is used
* `query` (string) - SQL query evaluated as the rule input

### RuleInputCustomNotebook
* `notebook` (string) - Fully-qualified notebook path, or a path relative to the configured base
  notebook directory
* `options` (object) - Optional arguments supplied when invoking the custom function

### RuleMitreMapping
* `sub_technique` (string) - Human-readable MITRE sub-technique name
* `sub_technique_id` (string) - MITRE sub-technique ID, such as `T1078.001`
* `tactic` (string) - Adversary objective during this attack, such as `Command and Control`
* `taxonomy` (string) - Structured framework that organizes adversary behaviors, such as
  `mitre-attack`
* `technique` (string) - Human-readable MITRE technique name
* `technique_id` (string) - MITRE technique ID, such as `T1078`

### RuleOutput
* `context` (object) - Key/value context to record from the input row. Values may contain SQL
  expressions in braces
* `default_context` (boolean) - If true, also record every column of the triggering input row. Custom
  context wins on key collision
* `event_time_column` (string) - Column in the rule output that holds the event timestamp
* `summary` (string) - Signal summary. May contain SQL expressions in braces that the
  server substitutes from the input row

### RuleSchedule
* `at_least_every` (string) - Defines the minimum cadence as a duration string. It may contain components
  for days, hours, minutes, and seconds (e.g. `5d`, `5d3h`, `3d2h5m12s`,
  `12m16s`)
* `continuous` (boolean) - Runs the rule continuously. Detection rules do not accept it: batch rules reject it, and
  streaming rules run continuously without it
* `cron_expression` (string) - Defines the schedule using a Quartz cron expression
* `enabled` (boolean) - Sets whether the job is scheduled to run. Defaults to true

### RuleTimeRange
* `extend_lookback` (string) - Extends the start of the execution time range by a compact duration with integer
  days, hours, minutes, and seconds in that order, such as `1h30m` or `15m`.
  Zero duration disables extension. Unset or empty uses a fixed `15m` at execution
  time. Only the start moves; the end is unchanged. Has no effect on batch SQL
  that does not use `IN_TIMERANGE`. Values must fit signed 64-bit microseconds