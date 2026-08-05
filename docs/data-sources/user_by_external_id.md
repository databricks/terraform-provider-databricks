---
subcategory: "Security"
---

# databricks_user_by_external_id Data Source

Resolves a [databricks_user](../resources/user.md) using the external ID assigned to it by the customer's Identity Provider (IdP).

-> This data source works with both the account-level and workspace-level provider.

-> **Get-or-create semantics.** This data source calls the Databricks `resolve-by-external-id` API, which creates the user in the account if no user with the given `external_id` exists yet. It is not a pure read: the first `terraform plan`/`apply` against a previously unseen `external_id` has the side effect of provisioning a new user. This API also requires the account to be onboarded to Automatic Identity Management (AIM); otherwise it returns an error.

## Example Usage

```hcl
data "databricks_user_by_external_id" "this" {
  external_id = "11111111-2222-3333-4444-555555555555"
}

resource "databricks_group_member" "member" {
  group_id  = databricks_group.this.id
  member_id = data.databricks_user_by_external_id.this.user_id
}
```

## Argument Reference

- `external_id` - (Required) The ID of the user in the customer's IdP (for example, the object ID assigned by Microsoft Entra ID).
- `api` - (Optional) Specifies whether to use account-level or workspace-level API. Valid values are `account` and `workspace`. When not set, the API level is inferred from the provider host.
- `provider_config` - (Optional) Configure the provider for management through account provider. This block consists of the following fields:
  - `workspace_id` - (Required) Workspace ID which the resource belongs to. This workspace must be part of the account which the provider is configured with.

## Attribute Reference

This data source exposes the following attributes:

- `user_id` - Internal ID of the user in Databricks.
- `account_id` - The parent account ID for the user in Databricks.
- `username` - Username/email of the user.
- `account_user_status` - The activity status of the user in the Databricks account.
- `full_name` - The full name of the user.
  - `given_name` - Given name of the user.
  - `family_name` - Family name of the user.

## Related Resources

The following resources are used in the same context:

- [**databricks_user**](../resources/user.md): Resource to manage individual users in Databricks.
- [**databricks_group_member**](../resources/group_member.md): Resource to manage group memberships by adding users to groups.
- [**databricks_users**](users.md): Data source to retrieve information about multiple users.
