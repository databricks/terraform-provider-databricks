---
subcategory: "Lakewatch"
---
# databricks_lakewatch_common_config Data Source
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
* `custom_content_pack_path` (string) - UC volume path holding user-authored custom content packs, e.g.
  `/Volumes/{catalog}/{schema}/{volume}`. User-provided with no default:
  when unset, custom content packs are unavailable (referencing one is
  rejected) and only built-in content packs apply
* `custom_presets_path` (string) - UC volume path holding user-authored custom presets, e.g.
  `/Volumes/{catalog}/{schema}/{volume}`. User-provided with no default:
  when unset, custom presets are unavailable and only built-in presets apply
* `etag` (string) - Opaque entity tag for optimistic concurrency. Returned on every read;
  pass it back as the Update request's `etag` to make the update
  conditional on the resource being unchanged
* `name` (string)
* `per_resource_permissions_enabled` (boolean) - Whether users may grant permissions on individual resources rather than
  only at the Folder level
* `update_time` (string)
* `updated_by` (Principal)
* `version` (integer) - Monotonic version, bumped on each update. Informational; use `etag` for
  optimistic concurrency

### Principal
* `display_name` (string) - Display name (user or SP name)
* `email` (string) - Email (regular users only)
* `id` (string) - User / service-principal ID
* `principal_type` (string) - Whether this principal is a regular user, service principal, or system principal. Possible values are: `REGULAR_USER`, `SERVICE_PRINCIPAL`, `SYSTEM_USER`