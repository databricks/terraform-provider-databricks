---
subcategory: "Agent Bricks"
---
# databricks_agent_kit_managed_memory_entries Data Source
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/agentkit)



## Example Usage


## Arguments
The following arguments are supported:
* `actor_id` (string, required) - Customer-provided identifier for the actor whose entries are listed
* `parent` (string, required) - Memory store whose entries are listed, in the form
  `memory-stores/{managed_memory_store_id}`
* `page_size` (integer, optional) - Maximum number of entries to return. The service may return fewer entries than requested.
  Defaults to 10; must be between 1 and 100
* `path_prefix` (string, optional) - Optional path prefix used to restrict entries within the actor partition
* `session_id` (string, optional) - Optional session identifier. When set, only entries with this exact `session_id` are
  returned. Omitted-session (cross-session) entries are not included. Ignored when path is set
* `view` (string, optional) - Controls whether each returned entry includes memory content. Defaults to `FULL`. Possible values are: `MANAGED_MEMORY_ENTRY_VIEW_BASIC`, `MANAGED_MEMORY_ENTRY_VIEW_FULL`
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.


## Attributes
This data source exports a single attribute, `managed_memory_entries`. It is a list of resources, each with the following attributes:
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