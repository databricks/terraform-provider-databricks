---
subcategory: "Identity and Access Management"
---
# databricks_workspace_iam_group_v2 Data Source
[![GA](https://img.shields.io/badge/Release_Stage-GA-green)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/iamv2)

Reads a single group by its resource name, `groups/{group_id}`, using a workspace-scoped provider.


## Example Usage
Example usage:

Look up a group by its internal ID, using a workspace-scoped provider.

```hcl
data "databricks_workspace_iam_group_v2" "this" {
  name = "groups/123456789"
}
```


## Arguments
The following arguments are supported:
* `name` (string, required) - The resource name of the group, in one of these formats:
  - Account APIs: `accounts/{account_id}/groups/{group_id}`
  - Workspace APIs: `groups/{group_id}`
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
The following attributes are exported:
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