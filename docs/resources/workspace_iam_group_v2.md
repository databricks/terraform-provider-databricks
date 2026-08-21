---
subcategory: "Identity and Access Management"
---
# databricks_workspace_iam_group_v2 Resource
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)

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
* `external_id` (string, optional) - ExternalId of the group in the customer's IdP
* `group_name` (string, optional) - Display name of the group
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
In addition to the above arguments, the following attributes are exported:
* `account_id` (string) - The parent account ID for group in Databricks
* `group_id` (string) - Internal group ID of the group in Databricks

## Import
As of Terraform v1.5, resources can be imported through configuration.
```hcl
import {
  id = "group_id"
  to = databricks_workspace_iam_group_v2.this
}
```

If you are using an older version of Terraform, import the resource using the `terraform import` command as follows:
```sh
terraform import databricks_workspace_iam_group_v2.this "group_id"
```