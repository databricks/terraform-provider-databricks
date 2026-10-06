---
subcategory: "Agent Bricks"
---
# databricks_mason_managed_memory_entry Data Source
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)



## Example Usage


## Arguments
The following arguments are supported:
* `name` (string, required) - Resource name in the form
  `memory-stores/{managed_memory_store_id}/entries/{managed_memory_entry_id}`
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
The following attributes are exported:
* `actor_id` (string) - Customer-provided identifier for the actor whose memory this entry represents
* `content` (string) - Optional free-form memory content
* `create_time` (string) - Time when the entry was created
* `description` (string) - Human-readable description of the memory entry
* `name` (string) - Resource name in the form
  `memory-stores/{managed_memory_store_id}/entries/{managed_memory_entry_id}`
* `path` (string) - Absolute, case-sensitive path identifying the entry within its actor and optional session.
  Paths must begin with `/` and must not contain empty, `.` or `..` segments
* `session_id` (string) - Optional identifier for the session associated with this memory entry. When omitted, the
  entry applies across the actor's sessions
* `source_type` (string) - Which writer created this entry. Caller sets this on Create; immutable after creation. Possible values are: `MANAGED_MEMORY_ENTRY_SOURCE_TYPE_AGENT`, `MANAGED_MEMORY_ENTRY_SOURCE_TYPE_DREAMER`
* `update_time` (string) - Time when the entry was last updated