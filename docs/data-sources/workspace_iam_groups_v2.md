---
subcategory: "Identity and Access Management"
---
# databricks_workspace_iam_groups_v2 Data Source
[![GA](https://img.shields.io/badge/Release_Stage-GA-green)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/iamv2)

Lists the groups visible to a workspace, using a workspace-scoped provider.


## Example Usage
Example usage:

List all groups visible to the workspace, using a workspace-scoped provider.

```hcl
data "databricks_workspace_iam_groups_v2" "this" {
}
```


## Arguments
The following arguments are supported:
* `filter` (string, optional) - Filters groups by `group_name` or `external_id`, for example `group_name eq "Data engineers"`.
  Put the value in quotes
* `page_size` (integer, optional) - The maximum number of groups to return, up to 1000. The default is 1000. A page can contain
  fewer groups
* `parent` (string, optional) - The resource name of the account to list groups in, in the format `accounts/{account_id}`.
  Defaults to the account that contains this workspace. Any other account is rejected
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.


## Attributes
This data source exports a single attribute, `groups`. It is a list of resources, each with the following attributes:
* `account_id` (string) - The ID of the account that contains the group
* `external_id` (string) - The ID of the group in your identity provider (IdP)
* `group_id` (string) - The Databricks ID of the group. It is a numeric string that Databricks assigns when the group
  is created, and it is the `{group_id}` segment of `name`
* `group_name` (string) - The display name of the group. It must be unique in the account, must not be blank, and must be
  shorter than 255 characters. Databricks trims surrounding whitespace. You can't use the name of a
  group that Databricks manages, such as `admins` or `users`
* `name` (string) - The resource name of the group, in one of these formats:
  - Account APIs: `accounts/{account_id}/groups/{group_id}`
  - Workspace APIs: `groups/{group_id}`
* `parent` (string) - The resource name of the account that contains the group, in the format
  `accounts/{account_id}`. Workspace APIs return the same format