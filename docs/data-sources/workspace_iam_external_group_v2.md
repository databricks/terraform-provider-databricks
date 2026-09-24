---
subcategory: "Identity and Access Management"
---
# databricks_workspace_iam_external_group_v2 Data Source
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/workspace/iamv2)

Retrieves an external group — a group that can be synced from your identity provider (IdP) — with the given external ID from the customer's IdP, scoped to a workspace. If the group does not exist in the account, it will be created. If the customer is not onboarded onto Automatic Identity Management (AIM), this returns an error.

The `name` uses the format `external-groups/{external_group_id}`, where `external_group_id` is the group's object ID in the IdP (for example, a Microsoft Entra ID object ID).

~> **Note** Reading this data source has a side effect: it resolves the group against the IdP and provisions the group in the account if it does not already exist. Provisioning happens at the account level; it does not assign the group to the workspace.

-> **Note** This data source can be used with an account-level or workspace-level provider. With an account-level provider, a `workspace_id` is required — set it in the `provider_config` block (or via the provider's `workspace_id` attribute). With a workspace-level provider, `workspace_id` is optional and defaults to the provider's workspace.


## Example Usage
Referring to an external group within a workspace, using an account-level provider with the target workspace selected via `provider_config`:

```hcl
data "databricks_workspace_iam_external_group_v2" "example" {
  name = "external-groups/11111111-2222-3333-4444-555555555555"
  provider_config = {
    workspace_id = "1234567890123456"
  }
}
```


## Arguments
The following arguments are supported:
* `name` (string, required) - The resource name of the external group. The format depends on the API that
  returned it:
  - Account-scoped: accounts/{account_id}/external-groups/{external_group_id}
  - Workspace-scoped: external-groups/{external_group_id}
* `provider_config` (ProviderConfig, optional) - Configure the provider for management through account provider.

### ProviderConfig
* `workspace_id` (string,optional) - Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attributes
The following attributes are exported:
* `account_id` (string) - The parent account ID, from Databricks
* `display_name` (string) - Display name of the group from the customer's IdP
* `external_group_id` (string) - The external ID of the group in the customer's IdP
* `internal_id` (string) - Internal groupId of the group in Databricks
* `name` (string) - The resource name of the external group. The format depends on the API that
  returned it:
  - Account-scoped: accounts/{account_id}/external-groups/{external_group_id}
  - Workspace-scoped: external-groups/{external_group_id}