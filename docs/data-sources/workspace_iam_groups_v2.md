---
subcategory: "Identity and Access Management"
---
# databricks_workspace_iam_groups_v2 Data Source
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)

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
* `filter` (string, optional) - Optional. Allows filtering groups by group name or external id
* `page_size` (integer, optional) - The maximum number of groups to return. The service may return fewer than this value
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.


## Attributes
This data source exports a single attribute, `groups`. It is a list of resources, each with the following attributes:
* `account_id` (string) - The parent account ID for group in Databricks
* `external_id` (string) - ExternalId of the group in the customer's IdP
* `group_id` (string) - Internal group ID of the group in Databricks
* `group_name` (string) - Display name of the group