---
subcategory: "Security"
---

# databricks_group_by_external_id Data Source

Resolves a [databricks_group](../resources/group.md) using the external ID assigned to it by the customer's Identity Provider (IdP).

-> This data source works with both the account-level and workspace-level provider.

-> **Get-or-create semantics.** This data source calls the Databricks `resolve-by-external-id` API, which creates the group in the account if no group with the given `external_id` exists yet. It is not a pure read: the first `terraform plan`/`apply` against a previously unseen `external_id` has the side effect of provisioning a new group. This API also requires the account to be onboarded to Automatic Identity Management (AIM); otherwise it returns an error.

## Example Usage

```hcl
data "databricks_group_by_external_id" "this" {
  external_id = "11111111-2222-3333-4444-555555555555"
}

resource "databricks_user" "this" {
  user_name = "me@example.com"
}

resource "databricks_group_member" "member" {
  group_id  = data.databricks_group_by_external_id.this.group_id
  member_id = databricks_user.this.id
}
```

## Argument Reference

- `external_id` - (Required) The ID of the group in the customer's IdP (for example, the object ID assigned by Microsoft Entra ID).
- `api` - (Optional) Specifies whether to use account-level or workspace-level API. Valid values are `account` and `workspace`. When not set, the API level is inferred from the provider host.
- `provider_config` - (Optional) Configure the provider for management through account provider. This block consists of the following fields:
  - `workspace_id` - (Required) Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attribute Reference

This data source exposes the following attributes:

- `group_id` - Internal ID of the group in Databricks.
- `account_id` - The parent account ID for the group in Databricks.
- `group_name` - Display name of the group.

## Related Resources

The following resources are used in the same context:

- [**databricks_group**](../resources/group.md): Resource to manage groups in Databricks.
- [**databricks_group_member**](../resources/group_member.md): Resource to manage group memberships by adding users to groups.
