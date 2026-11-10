---
subcategory: "Identity and Access Management"
---
# databricks_workspace_iam_users_v2 Data Source
[![GA](https://img.shields.io/badge/Release_Stage-GA-green)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/iamv2)

Lists the users visible to a workspace, using a workspace-scoped provider.


## Example Usage
Example usage:

List all users visible to the workspace, using a workspace-scoped provider.

```hcl
data "databricks_workspace_iam_users_v2" "this" {
}
```


## Arguments
The following arguments are supported:
* `filter` (string, optional) - Filters users by `username` or `external_id`, for example
  `username eq "someone@example.com"`. Put the value in quotes
* `page_size` (integer, optional) - The maximum number of users to return, up to 1000. The default is 1000. A page can contain
  fewer users
* `parent` (string, optional) - The resource name of the account to list users in, in the format `accounts/{account_id}`.
  Defaults to the account that contains this workspace. Any other account is rejected
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.


## Attributes
This data source exports a single attribute, `users`. It is a list of resources, each with the following attributes:
* `account_id` (string) - The ID of the account that contains the user
* `account_user_status` (string) - Whether the user is active in the account. New users are active unless you set this to
  `INACTIVE`. Possible values are: `ACTIVE`, `INACTIVE`
* `external_id` (string) - The ID of the user in your identity provider (IdP)
* `full_name` (UserFullName) - The user's name
* `name` (string) - The resource name of the user, in one of these formats:
  - Account APIs: `accounts/{account_id}/users/{user_id}`
  - Workspace APIs: `users/{user_id}`
* `parent` (string) - The resource name of the account that contains the user, in the format
  `accounts/{account_id}`. Workspace APIs return the same format
* `user_id` (string) - The Databricks ID of the user
* `username` (string) - The user's username, which is their email address. It can't be changed after the user is
  created

### UserFullName
* `family_name` (string) - The user's family (last) name
* `given_name` (string) - The user's given (first) name