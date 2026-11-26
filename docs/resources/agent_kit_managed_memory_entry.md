---
subcategory: "Agent Bricks"
---
# databricks_agent_kit_managed_memory_entry Resource
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/agentkit)



## Example Usage


## Arguments
The following arguments are supported:
* `actor_id` (string, required) - Customer-provided identifier for the actor whose memory this entry represents
* `parent` (string, required) - Memory store that will contain the entry, in the form
  `memory-stores/{managed_memory_store_id}`
* `path` (string, required) - Absolute, case-sensitive path identifying the entry within its actor and, until the
  deprecated `session_id` is ignored, its optional session.
  Paths must begin with `/` and must not contain empty, `.` or `..` segments
* `content` (string, optional) - Optional free-form memory content
* `description` (string, optional) - Human-readable description of the memory entry
* `managed_memory_entry_id` (string, optional) - Caller-selected memory entry ID. It must contain 1-255 letters, digits, hyphens, or
  underscores, and it becomes the last segment of the entry's `name`. The service generates a
  UUID when it is omitted
* `session_id` (string, optional, deprecated) - Deprecated: memory entries will be identified by `(actor_id, path)`, and the service will
  ignore this field on writes and omit it from responses. Until then, an omitted value means the
  entry applies across the actor's sessions
* `source_type` (string, optional) - Which writer created this entry. Caller sets this on Create; immutable after creation.
  The server may store an omitted or UNSPECIFIED value as AGENT. Possible values are: `MANAGED_MEMORY_ENTRY_SOURCE_TYPE_AGENT`, `MANAGED_MEMORY_ENTRY_SOURCE_TYPE_DREAMER`
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
In addition to the above arguments, the following attributes are exported:
* `create_time` (string) - Time when the entry was created
* `name` (string) - Resource name in the form
  `memory-stores/{managed_memory_store_id}/entries/{managed_memory_entry_id}`.
  Required for updates; assigned by the service on creation
* `update_time` (string) - Time when the entry was last updated

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "name"
  to = databricks_agent_kit_managed_memory_entry.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_agent_kit_managed_memory_entry.this "name"
```