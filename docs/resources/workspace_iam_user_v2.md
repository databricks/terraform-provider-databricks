---
subcategory: "Identity and Access Management"
---
# databricks_workspace_iam_user_v2 Resource
[![GA](https://img.shields.io/badge/Release_Stage-GA-green)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/iamv2)

Manages a user using a workspace-scoped provider. Creating the resource provisions the user; deleting it removes the user. `username` is immutable — changing it replaces the resource.

When Automatic Identity Management (AIM) is enabled for the account, this resource manages local users only: a user cannot be created with `external_id` set. Users that already have an `external_id` can still be read through this resource, but only their `external_id` may be updated — their other fields are sourced from the identity provider.


## Example Usage
Example usage:

Creates a user using a workspace-scoped provider. When AIM is enabled, the user cannot be created with `external_id` set.

```hcl
resource "databricks_workspace_iam_user_v2" "this" {
  username            = "jane.doe@example.com"
  account_user_status = "ACTIVE"

  full_name = {
    given_name  = "Jane"
    family_name = "Doe"
  }
}
```


## Arguments
The following arguments are supported:
* `account_user_status` (string, required) - Whether the user is active in the account. New users are active unless you set this to
  `INACTIVE`. Possible values are: `ACTIVE`, `INACTIVE`
* `full_name` (UserFullName, required) - The user's name
* `username` (string, required) - The user's username, which is their email address. It can't be changed after the user is
  created
* `external_id` (string, optional) - The ID of the user in your identity provider (IdP)
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

### UserFullName
* `family_name` (string, optional) - The user's family (last) name
* `given_name` (string, optional) - The user's given (first) name

## Attributes
In addition to the above arguments, the following attributes are exported:
* `account_id` (string) - The ID of the account that contains the user
* `name` (string) - The resource name of the user, in one of these formats:
  - Account APIs: `accounts/{account_id}/users/{user_id}`
  - Workspace APIs: `users/{user_id}`
* `parent` (string) - The resource name of the account that contains the user, in the format
  `accounts/{account_id}`. Workspace APIs return the same format
* `user_id` (string) - The Databricks ID of the user

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "name"
  to = databricks_workspace_iam_user_v2.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_workspace_iam_user_v2.this "name"
```