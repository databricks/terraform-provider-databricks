---
subcategory: "Agent Bricks"
---
# databricks_mason_session Data Source
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)



## Example Usage


## Arguments
The following arguments are supported:
* `name` (string, required) - Resource name in the form `session-stores/{session_store_id}/sessions/{session_id}`
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
The following attributes are exported:
* `actor_id` (string) - Opaque caller-provided identifier for the application actor associated with the session.
  
  This is application data and has no Databricks authentication or authorization semantics.
  Use the same value as the Managed Memory Entry `actor_id` when storing memories associated
  with this actor. Every session must set it. A child session must use the same value as its parent
* `create_time` (string) - Time when the session was created
* `last_activity_time` (string) - Time when the session's item history was last mutated
* `metadata` (object) - Mutable caller-defined string labels
* `name` (string) - Resource name in the form `session-stores/{session_store_id}/sessions/{session_id}`
* `parent_session_id` (string) - Immediate parent session ID. Set only at creation for child sessions, immutable thereafter, and
  restricted to the same store
* `root_session_id` (string) - Top-level session ID in the spawn tree. This equals `session_id` for a root or fork and is
  inherited transitively by child sessions
* `session_id` (string) - Unique session ID. The service generates a UUID unless the caller supplies
  `CreateSessionRequest.session_id`
* `update_time` (string) - Time when session resource fields last changed