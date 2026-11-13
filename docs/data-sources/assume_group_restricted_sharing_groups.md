---
subcategory: "Identity and Access Management"
---
# databricks_assume_group_restricted_sharing_groups Data Source
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)



## Example Usage


## Arguments
The following arguments are supported:
* `page_size` (integer, optional) - The maximum number of groups to return. The service may return fewer.
  If unspecified or 0, at most 100 groups are returned. Values above 100 are
  coerced to 100. Negative values are rejected
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.


## Attributes
This data source exports a single attribute, `restricted_sharing_groups`. It is a list of resources, each with the following attributes:
* `group_display_name` (string) - The display name of the restricted workspace group
* `name` (string) - The resource name of the restricted sharing group.
  Format: `restricted-sharing-groups/{restricted_sharing_group_id}`
  Example: `restricted-sharing-groups/123`
* `restricted_sharing_group_id` (string) - The id of the restricted workspace group: the numeric id of an existing and
  non-system workspace group