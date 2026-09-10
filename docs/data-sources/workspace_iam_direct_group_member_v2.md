---
subcategory: "Identity and Access Management"
---
# databricks_workspace_iam_direct_group_member_v2 Data Source
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/iamv2)

Reads a single direct group membership, scoped to a workspace, identified by group and principal ID.


## Example Usage
Example usage:

Look up a single direct group member by group and principal ID, using a
workspace-scoped provider.

```hcl
data "databricks_workspace_iam_direct_group_member_v2" "this" {
  group_id     = 123456789
  principal_id = 987654321
}
```


## Arguments
The following arguments are supported:
* `group_id` (integer, required) - The internal ID of the group this member belongs to
* `principal_id` (integer, required) - Internal ID of the principal in Databricks
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
The following attributes are exported:
* `display_name` (string) - Display name of the principal
* `external_id` (string) - The external ID of the principal in Databricks
* `group_id` (integer) - The internal ID of the group this member belongs to
* `membership_source` (string) - The source of group membership (internal or from identity provider). Possible values are: `IDENTITY_PROVIDER`, `INTERNAL`
* `principal_id` (integer) - Internal ID of the principal in Databricks
* `principal_type` (string) - The type of the principal (user/service principal/group). Possible values are: `GROUP`, `SERVICE_PRINCIPAL`, `USER`