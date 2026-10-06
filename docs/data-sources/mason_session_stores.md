---
subcategory: "Agent Bricks"
---
# databricks_mason_session_stores Data Source
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)



## Example Usage


## Arguments
The following arguments are supported:
* `page_size` (integer, optional) - Maximum number of session stores to return. Defaults to 10; must be between 1 and 100
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.


## Attributes
This data source exports a single attribute, `session_stores`. It is a list of resources, each with the following attributes:
* `create_time` (string) - Time when the store was created
* `creator_user_id` (string) - Workspace-local user ID of the authenticated principal that created the store. This is
  immutable server-set attribution and does not grant access; authorization is evaluated from
  the authenticated request context
* `description` (string) - Human-readable description of the session store
* `metadata` (object) - Mutable caller-defined string labels
* `name` (string) - Resource name in the form `session-stores/{session_store_id}`
* `update_time` (string) - Time when the store was last updated