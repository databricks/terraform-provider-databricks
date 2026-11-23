---
subcategory: "Identity and Access Management"
---
# databricks_workspace_iam_group_v2 Resource
[![GA](https://img.shields.io/badge/Release_Stage-GA-green)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/iamv2)

Manages a group using a workspace-scoped provider. Creating the resource creates the group; deleting it removes the group.

When Automatic Identity Management (AIM) is enabled for the account, this resource manages local groups only: a group cannot be created with `external_id` set. Groups that already have an `external_id` can still be read through this resource, but only their `external_id` may be updated — their other fields are sourced from the identity provider.


## Example Usage
Example usage:

Creates a group using a workspace-scoped provider. When AIM is enabled, the group cannot be created with `external_id` set.

```hcl
resource "databricks_workspace_iam_group_v2" "this" {
  group_name = "data-engineers"
}
```


## Arguments
The following arguments are supported:
* `external_id` (string, optional) - The ID of the group in your identity provider (IdP)
* `group_name` (string, optional) - The display name of the group. It must be unique in the account, must not be blank, and must be
  shorter than 255 characters. Databricks trims surrounding whitespace. You can't use the name of a
  group that Databricks manages, such as `admins` or `users`
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
In addition to the above arguments, the following attributes are exported:
* `account_id` (string) - The ID of the account that contains the group
* `group_id` (string) - The Databricks ID of the group. It is a numeric string that Databricks assigns when the group
  is created, and it is the `{group_id}` segment of `name`
* `name` (string) - The resource name of the group, in one of these formats:
  - Account APIs: `accounts/{account_id}/groups/{group_id}`
  - Workspace APIs: `groups/{group_id}`
* `parent` (string) - The resource name of the account that contains the group, in the format
  `accounts/{account_id}`. Workspace APIs return the same format

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "name"
  to = databricks_workspace_iam_group_v2.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_workspace_iam_group_v2.this "name"
```