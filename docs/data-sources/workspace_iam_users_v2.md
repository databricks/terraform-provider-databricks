---
subcategory: "Identity and Access Management"
---
# databricks_workspace_iam_users_v2 Data Source
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)

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
* `filter` (string, optional) - Optional. Allows filtering users by username or external id
* `page_size` (integer, optional) - The maximum number of users to return. The service may return fewer than this value
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.


## Attributes
This data source exports a single attribute, `users`. It is a list of resources, each with the following attributes:
* `account_id` (string) - The accountId parent of the user in Databricks
* `account_user_status` (string) - The activity status of a user in a Databricks account. Possible values are: `ACTIVE`, `INACTIVE`
* `external_id` (string) - ExternalId of the user in the customer's IdP
* `full_name` (UserFullName)
* `user_id` (string) - Internal userId of the user in Databricks
* `username` (string) - Username/email of the user

### UserFullName
* `family_name` (string)
* `given_name` (string)