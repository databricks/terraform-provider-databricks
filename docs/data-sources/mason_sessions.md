---
subcategory: "Agent Bricks"
---
# databricks_mason_sessions Data Source
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)



## Example Usage


## Arguments
The following arguments are supported:
* `parent` (string, required) - Resource name of the containing session store, in the form
  `session-stores/{session_store_id}`
* `filter` (string, optional) - Filter expression. Supported fields include `actor_id` and `metadata`; for example,
  `actor_id = "support-customer-123"`
* `order_by` (string, optional) - Sort order. Defaults to `last_activity_time desc`. Page-token continuation is exactly-once when
  ordering by `create_time` (immutable); ordering by `last_activity_time` is best-effort, because
  that value changes as a session gains activity, so a session updated between page requests may
  be repeated or skipped. To enumerate every session exactly once, order by `create_time`
* `page_size` (integer, optional) - Maximum number of sessions to return. Defaults to 10; must be between 1 and 100
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.


## Attributes
This data source exports a single attribute, `sessions`. It is a list of resources, each with the following attributes:
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