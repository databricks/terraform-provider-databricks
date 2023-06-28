---
subcategory: "Security"
---

# databricks_rule_set Resource

This resource allows you to manage access rules on Databricks account level resources. For convenience we allow accessing this resource through the Databricks account and workspace.

To configure rule sets through the Databricks account, the provider must be configured with `host = "https://accounts.cloud.databricks.com"` on AWS deployments or `host = "https://accounts.azuredatabricks.net"` and authenticate using [AAD tokens](https://registry.terraform.io/providers/databricks/databricks/latest/docs#special-configurations-for-azure) on Azure deployments

## Example usage

Rule set management through AWS Databricks workspace:

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

resource "databricks_rule_set" "automation_sp_rule_set" {
  name = "accounts/${local.account_id}/servicePrincipals/${databricks_service_principal.automation_sp.application_id}/ruleSets/default"

  grant_rules {
    principals = ["groups/${data.databricks_group.ds.display_name}"]
    role       = "roles/servicePrincipal.user"
  }
}
```

Rule set management through AWS Databricks account:

```hcl
// initialize provider at account-level
provider "databricks" {
  alias      = "mws"
  host       = "https://accounts.cloud.databricks.com"
  account_id = "00000000-0000-0000-0000-000000000000"
  username   = var.databricks_account_username
  password   = var.databricks_account_password
}

// account level group creation
resource "databricks_group" "ds" {
  display_name = "Data Science"
}

resource "databricks_service_principal" "automation_sp" {
  display_name = "SP_FOR_AUTOMATION"
}

resource "databricks_rule_set" "automation_sp_rule_set" {
  name = "accounts/${local.account_id}/servicePrincipals/${databricks_service_principal.automation_sp.application_id}/ruleSets/default"

  grant_rules {
    principals = ["groups/${databricks_group.ds.display_name}"]
    role       = "roles/servicePrincipal.user"
  }
}
```

Rule set management through Azure Databricks account:

```hcl
// initialize provider at Azure account-level
provider "databricks" {
  alias      = "azure_account"
  host       = "https://accounts.azuredatabricks.net"
  account_id = "00000000-0000-0000-0000-000000000000"
  auth_type  = "azure-cli"
}

// account level group creation
resource "databricks_group" "ds" {
  display_name = "Data Science"
}

resource "databricks_service_principal" "automation_sp" {
  application_id = "00000000-0000-0000-0000-000000000000"
  display_name   = "SP_FOR_AUTOMATION"
}

resource "databricks_rule_set" "automation_sp_rule_set" {
  name = "accounts/${local.account_id}/servicePrincipals/${databricks_service_principal.automation_sp.application_id}/ruleSets/default"

  grant_rules {
    principals = ["groups/${databricks_group.ds.display_name}"]
    role       = "roles/servicePrincipal.user"
  }
}
```

## Argument Reference

Rule set name and grant rules block argument are required.

!> **Warning** Name uniquely identifies a rule set resource. Ensure all the grant_rules blocks for a rule set name are present in one rule_set resource block. Otherwise, after applying changes, users might lose their role assignment even if that was not intended.

### name

Unique identifier of a rule set. Currently, only default rule sets are supported. The following rule set formats are supported:
* accounts/{account_id}/servicePrincipals/{service_principal_application_id}/ruleSets/default

### grant_rules block

One or more `grant_rules` blocks are required to actually set access rules.

```hcl
grant_rules {
  principals = [
    "groups/{databricks_group.ds.display_name}"
  ]
  role = "roles/servicePrincipal.user"
}
```

Arguments of the `grant_rules` block are:

- `role` - (Required) Role to be granted. The following roles are supported:
  * roles/servicePrincipal.manager - Manager of a service principal.
  * roles/servicePrincipal.user - User of a service principal.
- `principals` - (Required) a list of principals who are granted a role. The following format is supported:
  * users/{username}
  * groups/{groupname}
  * servicePrincipals/{applicationId}

## Related Resources

The following resources are often used in the same context:

* [databricks_group](group.md)
* [databricks_user](user.md)
* [databricks_service_principal](service_principal.md)
