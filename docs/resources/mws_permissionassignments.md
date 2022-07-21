---
subcategory: "Unity Catalog"
---
# databricks_mws_permissionassignments Resource

TBA

## Example Usage

Adding account group

```hcl
provider "databricks" {
    account_id = <databricks account id>
}

resource "databricks_group" "data_eng" {
  display_name               = "Data Engineering"
}

resource "databricks_mws_permissionassignments" "add_admin_group" {
    workspace_id = databricks_mws_workspaces.this.workspace_id
    principal_id = databricks_group.data_eng.id
    permissions = ["ADMIN"]
}
```

Adding account user

```hcl
provider "databricks" {
    account_id = <databricks account id>
}

resource "databricks_user" "me" {
  user_name = "me@example.com"
}

resource "databricks_mws_permissionassignments" "add_user" {
    workspace_id = databricks_mws_workspaces.this.workspace_id
    principal_id = databricks_user.me.id
    permissions = ["USER"]
}
```

Adding account service principal

```hcl
provider "databricks" {
    account_id = <databricks account id>
}

resource "databricks_service_principal" "sp" {
  display_name = "Automation-only SP"
}

resource "databricks_mws_permissionassignments" "add_admin_spn" {
    workspace_id = databricks_mws_workspaces.this.workspace_id
    principal_id = databricks_service_principal.sp.id
    permissions = ["ADMIN"]
}
```