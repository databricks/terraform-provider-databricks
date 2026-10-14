---
subcategory: "Agent Bricks"
---
# databricks_mason_managed_memory_stores Data Source
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)



## Example Usage


## Arguments
The following arguments are supported:
* `page_size` (integer, optional) - Maximum number of stores to return. The service may return fewer stores than requested.
  Defaults to 10; must be between 1 and 100
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.


## Attributes
This data source exports a single attribute, `managed_memory_stores`. It is a list of resources, each with the following attributes:
* `create_time` (string) - Time when the store was created
* `creator_user_id` (string) - Workspace-local user ID of the authenticated principal that created the store. This is
  immutable server-set attribution and does not grant access; authorization is evaluated from
  the authenticated request context
* `description` (string) - Human-readable description of the memory store
* `display_name` (string, deprecated) - Deprecated compatibility alias for the caller-provided managed memory store ID. Canonical
  clients provide the ID through `CreateMemoryStoreRequest.managed_memory_store_id` and use
  `name` as the resource identifier
* `name` (string) - Resource name in the form `memory-stores/{managed_memory_store_id}`
* `owner_user_id` (string, deprecated) - Deprecated alias for `creator_user_id`. This identifies the original creator, not a
  transferable owner. Use `creator_user_id` instead
* `storage_backend` (StorageBackend) - Service-managed storage backing this memory store
* `update_time` (string) - Time when the store was last updated
* `workspace_id` (integer) - Workspace that owns the memory store

### StorageBackend
* `backend_id` (string) - Backend-specific identifier. For Lakebase, this is the project ID
* `backend_type` (string) - Type of the storage backend. Possible values are: `STORAGE_BACKEND_TYPE_LAKEBASE`