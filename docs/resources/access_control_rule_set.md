---
subcategory: "Security"
---

# databricks_access_control_rule_set Resource

-> **Note** This resource can be used with an account or workspace-level provider.

This resource allows you to manage access rules on Databricks account level resources. For convenience we allow accessing this resource through the Databricks account and workspace.

-> **Note** Currently, we only support managing access rules on service principal, group and account resources through `databricks_access_control_rule_set`.

-> **Warning** `databricks_access_control_rule_set` cannot be used to manage access rules for resources supported by [databricks_permissions](permissions.md). Refer to its documentation for more information.

## Service principal rule set usage

Through a Databricks workspace:

```hcl
locals {
  account_id = "00000000-0000-0000-0000-000000000000"
}

// account level group
data "databricks_group" "ds" {
  display_name = "Data Science"
}

resource "databricks_service_principal" "automation_sp" {
  display_name = "SP_FOR_AUTOMATION"
}

resource "databricks_access_control_rule_set" "automation_sp_rule_set" {
  name = "accounts/${local.account_id}/servicePrincipals/${databricks_service_principal.automation_sp.application_id}/ruleSets/default"

  grant_rules {
    principals = [data.databricks_group.ds.acl_principal_id]
    role       = "roles/servicePrincipal.user"
  }
}
```

Through AWS Databricks account:

```hcl
locals {
  account_id = "00000000-0000-0000-0000-000000000000"
}

// initialize provider at account-level
provider "databricks" {
  host          = "https://accounts.cloud.databricks.com"
  account_id    = local.account_id
  client_id     = var.client_id
  client_secret = var.client_secret
}

// account level group creation
resource "databricks_group" "ds" {
  display_name = "Data Science"
}

resource "databricks_service_principal" "automation_sp" {
  display_name = "SP_FOR_AUTOMATION"
}

resource "databricks_access_control_rule_set" "automation_sp_rule_set" {
  name = "accounts/${local.account_id}/servicePrincipals/${databricks_service_principal.automation_sp.application_id}/ruleSets/default"

  grant_rules {
    principals = [databricks_group.ds.acl_principal_id]
    role       = "roles/servicePrincipal.user"
  }
}
```

Through Azure Databricks account:

```hcl
locals {
  account_id = "00000000-0000-0000-0000-000000000000"
}

// initialize provider at Azure account-level
provider "databricks" {
  host       = "https://accounts.azuredatabricks.net"
  account_id = local.account_id
}

// account level group creation
resource "databricks_group" "ds" {
  display_name = "Data Science"
}

resource "databricks_service_principal" "automation_sp" {
  application_id = "00000000-0000-0000-0000-000000000000"
  display_name   = "SP_FOR_AUTOMATION"
}

resource "databricks_access_control_rule_set" "automation_sp_rule_set" {
  name = "accounts/${local.account_id}/servicePrincipals/${databricks_service_principal.automation_sp.application_id}/ruleSets/default"

  grant_rules {
    principals = [databricks_group.ds.acl_principal_id]
    role       = "roles/servicePrincipal.user"
  }
}
```

Through GCP Databricks account:

```hcl
locals {
  account_id = "00000000-0000-0000-0000-000000000000"
}

// initialize provider at account-level
provider "databricks" {
  host       = "https://accounts.gcp.databricks.com"
  account_id = local.account_id
}

// account level group creation
resource "databricks_group" "ds" {
  display_name = "Data Science"
}

resource "databricks_service_principal" "automation_sp" {
  display_name = "SP_FOR_AUTOMATION"
}

resource "databricks_access_control_rule_set" "automation_sp_rule_set" {
  name = "accounts/${local.account_id}/servicePrincipals/${databricks_service_principal.automation_sp.application_id}/ruleSets/default"

  grant_rules {
    principals = [databricks_group.ds.acl_principal_id]
    role       = "roles/servicePrincipal.user"
  }
}
```

## Group rule set usage

Refer to the appropriate provider configuration as shown in the examples for service principal rule set.

```hcl
locals {
  account_id = "00000000-0000-0000-0000-000000000000"
}

// account level group
data "databricks_group" "ds" {
  display_name = "Data Science"
}

data "databricks_user" "john" {
  user_name = "john.doe@example.com"
}

resource "databricks_access_control_rule_set" "ds_group_rule_set" {
  name = "accounts/${local.account_id}/groups/${databricks_group.ds.id}/ruleSets/default"

  grant_rules {
    principals = [data.databricks_user.john.acl_principal_id]
    role       = "roles/group.manager"
  }
}
```

## Account rule set usage

Refer to the appropriate provider configuration as shown in the examples for service principal rule set.

```hcl
locals {
  account_id = "00000000-0000-0000-0000-000000000000"
}

// account level group
data "databricks_group" "ds" {
  display_name = "Data Science"
}

// account level group
data "databricks_group" "marketplace_admins" {
  display_name = "Marketplace Admins"
}

data "databricks_user" "john" {
  user_name = "john.doe@example.com"
}

resource "databricks_access_control_rule_set" "account_rule_set" {
  name = "accounts/${local.account_id}/ruleSets/default"

  // user john is manager for all groups in the account
  grant_rules {
    principals = [data.databricks_user.john.acl_principal_id]
    role       = "roles/group.manager"
  }

  // group data science is manager for all service principals in the account
  grant_rules {
    principals = [data.databricks_group.ds.acl_principal_id]
    role       = "roles/servicePrincipal.manager"
  }

  grant_rules {
    principals = [data.databricks_group.marketplace_admins.acl_principal_id]
    role       = "roles/marketplace.admin"
  }
}
```

## Argument Reference

* `name` - (Required) Unique identifier of a rule set. The name determines the resource to which the rule set applies. Currently, only default rule sets are supported. The following rule set formats are supported:
  * `accounts/{account_id}/servicePrincipals/{service_principal_application_id}/ruleSets/default`
  * `accounts/{account_id}/groups/{group_id}/ruleSets/default`
  * `accounts/{account_id}/ruleSets/default`

* `grant_rules` - (Required) The access control rules to be granted by this rule set, consisting of a set of principals and roles to be granted to them.

!> **Warning** Name uniquely identifies a rule set resource. Ensure all the grant_rules blocks for a rule set name are present in one `databricks_access_control_rule_set` resource block. Otherwise, after applying changes, users might lose their role assignment even if that was not intended.

### grant_rules

One or more `grant_rules` blocks are required to actually set access rules.

```hcl
grant_rules {
  principals = [
    databricks_group.ds.acl_principal_id
  ]
  role = "roles/servicePrincipal.user"
}
```

Arguments of the `grant_rules` block are:

* `role` - (Required) Role to be granted. The supported roles are listed below. For more information about these roles, refer to [service principal roles](https://docs.databricks.com/security/auth-authz/access-control/service-principal-acl.html#service-principal-roles), [group roles](https://docs.databricks.com/en/administration-guide/users-groups/groups.html#manage-roles-on-an-account-group-using-the-workspace-admin-settings-page) or [marketplace roles](https://docs.databricks.com/en/marketplace/get-started-provider.html#assign-the-marketplace-admin-role).
  * `roles/servicePrincipal.manager` - Manager of a service principal.
  * `roles/servicePrincipal.user` - User of a service principal.
  * `roles/group.manager` - Manager of a group.
  * `roles/marketplace.admin` - Admin of marketplace.
* `principals` - (Required) a list of principals who are granted a role. The following format is supported:
  * `users/{username}` (also exposed as `acl_principal_id` attribute of `databricks_user` resource).
  * `groups/{groupname}` (also exposed as `acl_principal_id` attribute of `databricks_group` resource).
  * `servicePrincipals/{applicationId}` (also exposed as `acl_principal_id` attribute of `databricks_service_principal` resource).

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the access control rule set - the same as `name`.

## Related Resources

The following resources are often used in the same context:

* [databricks_group](group.md)
* [databricks_user](user.md)
* [databricks_service_principal](service_principal.md)
