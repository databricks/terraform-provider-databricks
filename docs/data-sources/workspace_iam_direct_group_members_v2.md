---
subcategory: "Identity and Access Management"
---
# databricks_workspace_iam_direct_group_members_v2 Data Source
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/iamv2)

Lists the direct members of a group, scoped to a workspace.


## Example Usage
Example usage:

List all direct members of a group, using a workspace-scoped provider.

```hcl
data "databricks_workspace_iam_direct_group_members_v2" "this" {
  group_id = 123456789
}
```


## Arguments
The following arguments are supported:
* `group_id` (integer, required) - Required. Internal ID of the group in Databricks whose direct members are being listed
* `page_size` (integer, optional) - The maximum number of members to return. The service may return fewer than this value.
  If not provided, defaults to 1000 (also the maximum allowed)
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.


## Attributes
This data source exports a single attribute, `direct_group_members`. It is a list of resources, each with the following attributes:
* `display_name` (string) - Display name of the principal
* `external_id` (string) - The external ID of the principal in Databricks
* `group_id` (integer) - The internal ID of the group this member belongs to
* `membership_source` (string) - The source of group membership (internal or from identity provider). Possible values are: `IDENTITY_PROVIDER`, `INTERNAL`
* `principal_id` (integer) - Internal ID of the principal in Databricks
* `principal_type` (string) - The type of the principal (user/service principal/group). Possible values are: `GROUP`, `SERVICE_PRINCIPAL`, `USER`