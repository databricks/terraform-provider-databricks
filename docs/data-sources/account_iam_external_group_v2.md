---
subcategory: "Identity and Access Management"
---
# databricks_account_iam_external_group_v2 Data Source
[![GA](https://img.shields.io/badge/Release_Stage-GA-green)](https://docs.databricks.com/aws/en/release-notes/release-types)

[API Documentation](https://docs.databricks.com/api/account/iamv2)

Retrieves an external group — a group that can be synced from your identity provider (IdP) — with the given external ID from the customer's IdP. If the group does not exist in the account, it will be created. If the customer is not onboarded onto Automatic Identity Management (AIM), this returns an error.

The `name` uses the format `accounts/{account_id}/external-groups/{external_group_id}`, where `external_group_id` is the group's object ID in the IdP (for example, a Microsoft Entra ID object ID).

~> **Note** Reading this data source has a side effect: it resolves the group against the IdP and provisions the group in the account if it does not already exist.

-> **Note** This data source can only be used with an account-level provider.


## Example Usage
Referring to an external group by its resource name:

```hcl
data "databricks_account_iam_external_group_v2" "example" {
  name = "accounts/00000000-0000-0000-0000-000000000000/external-groups/11111111-2222-3333-4444-555555555555"
}
```


## Arguments
The following arguments are supported:
* `name` (string, required) - The resource name of the external group, in one of these formats:
  - Account APIs: `accounts/{account_id}/external-groups/{external_group_id}`
  - Workspace APIs: `external-groups/{external_group_id}`

## Attributes
The following attributes are exported:
* `account_id` (string) - The ID of the account that contains the group
* `display_name` (string) - The name of the group, from your IdP. It's the same as the group's `group_name`
* `external_group_id` (string) - The ID of the group in your IdP
* `internal_id` (string) - The Databricks ID of the group. Use it as the `group_id` in the group APIs
* `name` (string) - The resource name of the external group, in one of these formats:
  - Account APIs: `accounts/{account_id}/external-groups/{external_group_id}`
  - Workspace APIs: `external-groups/{external_group_id}`