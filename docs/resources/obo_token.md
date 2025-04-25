---
subcategory: "Security"
---
# databricks_obo_token Resource

This resource creates [On-Behalf-Of tokens](https://docs.databricks.com/administration-guide/users-groups/service-principals.html#manage-personal-access-tokens-for-a-service-principal) for a [databricks_service_principal](service_principal.md) in Databricks workspaces on AWS and GCP.  In general it's best to use OAuth authentication using client ID and secret, and use this resource mostly for integrations that doesn't support OAuth.

-> This resource can only be used with a workspace-level provider!

~> To create On-Behalf-Of token for Azure Service Principal, configure Terraform provider to use Azure service principal's client ID and secret, and use `databricks_token` resource to create a personal access token.

## Example Usage

Creating a token for a narrowly-scoped service principal, that would be the only one (besides admins) allowed to use PAT token in this given workspace, keeping your automated deployment highly secure.

-> A given declaration of `databricks_permissions.token_usage` would OVERWRITE permissions to use PAT tokens from any existing groups with token usage permissions such as the `users` group. To avoid this, be sure to include any desired groups in additional `access_control` blocks in the Terraform configuration file.

```hcl
resource "databricks_service_principal" "this" {
  display_name = "Automation-only SP"
}

resource "databricks_permissions" "token_usage" {
  authorization = "tokens"
  access_control {
    service_principal_name = databricks_service_principal.this.application_id
    permission_level       = "CAN_USE"
  }
}

resource "databricks_obo_token" "this" {
  depends_on       = [databricks_permissions.token_usage]
  application_id   = databricks_service_principal.this.application_id
  comment          = "PAT on behalf of ${databricks_service_principal.this.display_name}"
  lifetime_seconds = 3600
}

output "obo" {
  value     = databricks_obo_token.this.token_value
  sensitive = true
}
```

Creating a token for a service principal with admin privileges

```hcl
resource "databricks_service_principal" "this" {
  display_name = "Terraform"
}

data "databricks_group" "admins" {
  display_name = "admins"
}

resource "databricks_group_member" "this" {
  group_id  = data.databricks_group.admins.id
  member_id = databricks_service_principal.this.id
}

resource "databricks_obo_token" "this" {
  depends_on       = [databricks_group_member.this]
  application_id   = databricks_service_principal.this.application_id
  comment          = "PAT on behalf of ${databricks_service_principal.this.display_name}"
  lifetime_seconds = 3600
}
```

## Argument Reference

The following arguments are required:

* `application_id` - Application ID of [databricks_service_principal](service_principal.md#application_id) to create a PAT token for.
* `lifetime_seconds` - (Integer, Optional) The number of seconds before the token expires. Token resource is re-created when it expires. If no lifetime is specified, the token remains valid indefinitely.
* `comment` - (String, Optional) Comment that describes the purpose of the token.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Canonical unique identifier for the token.
* `token_value` - **Sensitive** value of the newly-created token.

## Import

!> Importing this resource is not currently supported.

## Related Resources

The following resources are often used in the same context:

* [End to end workspace management](../guides/workspace-management.md) guide.
* [databricks_group](../data-sources/group.md) data to retrieve information about [databricks_group](group.md) members, entitlements and instance profiles.
* [databricks_group_member](group_member.md) to attach [users](user.md) and [groups](group.md) as group members.
* [databricks_permissions](permissions.md) to manage [access control](https://docs.databricks.com/security/access-control/index.html) in Databricks workspace.
* [databricks_service_principal](service_principal.md) to manage [Service Principals](https://docs.databricks.com/administration-guide/users-groups/service-principals.html) that could be added to [databricks_group](group.md) within workspace.
* [databricks_sql_permissions](sql_permissions.md) to manage data object access control lists in Databricks workspaces for things like tables, views, databases, and [more](https://docs.databricks.com/security/access-control/table-acls/object-privileges.html).
