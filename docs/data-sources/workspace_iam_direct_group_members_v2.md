---
subcategory: "Identity and Access Management"
---
# databricks_workspace_iam_direct_group_members_v2 Data Source
[![GA](https://img.shields.io/badge/Release_Stage-GA-green)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/iamv2)

Lists the direct members of a group, scoped to a workspace.


## Example Usage
Example usage:

List all direct members of a group, using a workspace-scoped provider.

```hcl
data "databricks_workspace_iam_direct_group_members_v2" "this" {
  parent = "groups/123456789"
}
```


## Arguments
The following arguments are supported:
* `parent` (string, required) - The resource name of the group, in the format `groups/{group_id}`
* `page_size` (integer, optional) - The maximum number of members to return, up to 1000. The default is 1000. A page can contain
  fewer members
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.


## Attributes
This data source exports a single attribute, `direct_group_members`. It is a list of resources, each with the following attributes:
* `display_name` (string) - The member's display name
* `external_id` (string) - The ID of the member in your identity provider (IdP)
* `group_id` (integer) - The Databricks ID of the group
* `membership_source` (string) - Where the membership comes from. Possible values are: `IDENTITY_PROVIDER`, `INTERNAL`
* `name` (string) - The resource name of the membership, in one of these formats:
  - Account APIs: `accounts/{account_id}/groups/{group_id}/direct-members/{principal_id}`
  - Workspace APIs: `groups/{group_id}/direct-members/{principal_id}`
* `parent` (string) - The resource name of the group, in one of these formats:
  - Account APIs: `accounts/{account_id}/groups/{group_id}`
  - Workspace APIs: `groups/{group_id}`
* `principal_id` (integer) - The Databricks ID of the member
* `principal_type` (string) - Whether the member is a user, service principal, or group. Possible values are: `GROUP`, `SERVICE_PRINCIPAL`, `USER`