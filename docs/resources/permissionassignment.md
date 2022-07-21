---
subcategory: "Unity Catalog"
---
# databricks_permissionassignment Resource

TBA

These resources are invoked in the workspace context.

## Example Usage

Adding workspace group

```hcl
resource "databricks_group" "data_eng" {
  display_name = "Data Engineering"
}

resource "databricks_permissionassignment" "add_admin_group" {
  principal_id = databricks_group.data_eng.id
  permissions  = ["ADMIN"]
}
```

Adding workspace user

```hcl
resource "databricks_user" "me" {
  user_name = "me@example.com"
}

resource "databricks_permissionassignment" "add_user" {
  principal_id = databricks_user.me.id
  permissions  = ["USER"]
}
```

Adding workspace service principal

```hcl
resource "databricks_service_principal" "sp" {
  display_name = "Automation-only SP"
}

resource "databricks_permissionassignment" "add_admin_spn" {
  principal_id = databricks_service_principal.sp.id
  permissions  = ["ADMIN"]
}
```
