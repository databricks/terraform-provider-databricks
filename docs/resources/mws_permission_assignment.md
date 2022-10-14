---
subcategory: "Security"
---
# databricks_mws_permission_assignment Resource

These resources are invoked in the account context. Provider must have `account_id` attribute configured.

## Example Usage

In account context, adding account-level group to a workspace:

```hcl
provider "databricks" {
  // <other properties>
  account_id = "<databricks account id>"
}

resource "databricks_group" "data_eng" {
  display_name = "Data Engineering"
}

resource "databricks_mws_permission_assignment" "add_admin_group" {
  workspace_id = databricks_mws_workspaces.this.workspace_id
  principal_id = databricks_group.data_eng.id
  permissions  = ["ADMIN"]
}
```

In account context, adding account-level user to a workspace:

```hcl
provider "databricks" {
  // <other properties>
  account_id = "<databricks account id>"
}

resource "databricks_user" "me" {
  user_name = "me@example.com"
}

resource "databricks_mws_permission_assignment" "add_user" {
  workspace_id = databricks_mws_workspaces.this.workspace_id
  principal_id = databricks_user.me.id
  permissions  = ["USER"]
}
```

In account context, adding account-level service principal to a workspace:

```hcl
provider "databricks" {
  // <other properties>
  account_id = "<databricks account id>"
}

resource "databricks_service_principal" "sp" {
  display_name = "Automation-only SP"
}

resource "databricks_mws_permission_assignment" "add_admin_spn" {
  workspace_id = databricks_mws_workspaces.this.workspace_id
  principal_id = databricks_service_principal.sp.id
  permissions  = ["ADMIN"]
}
```
