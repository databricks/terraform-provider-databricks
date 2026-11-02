---
subcategory: "Identity and Access Management"
---
# databricks_workspace_iam_group_v2 Data Source
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

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
* `name` (string, required) - The resource name of the group. The format depends on the API that returned it:
  - Account-scoped: accounts/{account_id}/groups/{group_id}
  - Workspace-scoped: groups/{group_id}
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
The following attributes are exported:
* `account_id` (string) - The parent account ID for group in Databricks
* `external_id` (string) - ExternalId of the group in the customer's IdP
* `group_id` (string) - Internal group ID of the group in Databricks
* `group_name` (string) - Display name of the group
* `name` (string) - The resource name of the group. The format depends on the API that returned it:
  - Account-scoped: accounts/{account_id}/groups/{group_id}
  - Workspace-scoped: groups/{group_id}
* `parent` (string) - The resource name of the account that contains the group, `accounts/{account_id}`. Both the
  account-scoped and workspace-scoped APIs return it