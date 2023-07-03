---
subcategory: "Security"
---

# databricks_access_control_rule_set Resource

This resource allows you to manage access rules on Databricks account level resources. For convenience we allow accessing this resource through the Databricks account and workspace.

-> **Note** Currently, we only support managing access rules on service principal resources through databricks_access_control_rule_set.

-> **Warning** databricks_access_control_rule_set cannot be used to manage access rules for resources supported by [databricks_permissions](permissions.md). Refer to its documentation for more information.

## Example usage

Rule set management through a Databricks workspace:

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

resource "databricks_access_control_rule_set" "automation_sp_rule_set" {
  name = "accounts/${provider.databricks.account_id}/servicePrincipals/${databricks_service_principal.automation_sp.application_id}/ruleSets/default"

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

resource "databricks_access_control_rule_set" "automation_sp_rule_set" {
  name = "accounts/${provider.databricks.account_id}/servicePrincipals/${databricks_service_principal.automation_sp.application_id}/ruleSets/default"

  grant_rules {
    principals = ["groups/${databricks_group.ds.display_name}"]
    role       = "roles/servicePrincipal.user"
  }
}
```

Rule set management through GCP Databricks account:

```hcl
// initialize provider at account-level
provider "databricks" {
  alias      = "mws"
  host       = "https://accounts.gcp.databricks.com"
  account_id = "00000000-0000-0000-0000-000000000000"
}

// account level group creation
resource "databricks_group" "ds" {
  display_name = "Data Science"
}

resource "databricks_service_principal" "automation_sp" {
  display_name = "SP_FOR_AUTOMATION"
}

resource "databricks_access_control_rule_set" "automation_sp_rule_set" {
  name = "accounts/${provider.databricks.account_id}/servicePrincipals/${databricks_service_principal.automation_sp.application_id}/ruleSets/default"

  grant_rules {
    principals = ["groups/${databricks_group.ds.display_name}"]
    role       = "roles/servicePrincipal.user"
  }
}
```

## Argument Reference

* `name` - (Required) Unique identifier of a rule set. The name determines the resource to which the rule set applies. Currently, only default rule sets are supported. The following rule set formats are supported:
  * accounts/{account_id}/servicePrincipals/{service_principal_application_id}/ruleSets/default

* `grant_rules` - (Required) The access control rules to be granted by this rule set, consisting of a set of principals and roles to be granted to them.

!> **Warning** Name uniquely identifies a rule set resource. Ensure all the grant_rules blocks for a rule set name are present in one databricks_access_control_rule_set resource block. Otherwise, after applying changes, users might lose their role assignment even if that was not intended.

### grant_rules

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
