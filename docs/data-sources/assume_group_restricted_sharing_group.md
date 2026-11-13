---
subcategory: "Identity and Access Management"
---
# databricks_assume_group_restricted_sharing_group Data Source
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)



## Example Usage


## Arguments
The following arguments are supported:
* `name` (string, required) - The resource name of the restricted sharing group.
  Format: `restricted-sharing-groups/{restricted_sharing_group_id}`
  Example: `restricted-sharing-groups/123`
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
The following attributes are exported:
* `group_display_name` (string) - The display name of the restricted workspace group
* `name` (string) - The resource name of the restricted sharing group.
  Format: `restricted-sharing-groups/{restricted_sharing_group_id}`
  Example: `restricted-sharing-groups/123`
* `restricted_sharing_group_id` (string) - The id of the restricted workspace group: the numeric id of an existing and
  non-system workspace group