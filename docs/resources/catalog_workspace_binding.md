---
subcategory: "Unity Catalog"
---
# databricks_catalog_workspace_binding Resource

If you use workspaces to isolate user data access, you may want to limit catalog access to specific workspaces in your account, also known as workspace-catalog binding

By default, Databricks assigns the catalog to all workspaces attached to the current metastore. By using `databricks_catalog_workspace_binding`, the catalog will be unassigned from all workspaces and only assigned explicitly using this resource. 

## Example Usage

```hcl
resource "databricks_catalog_workspace_binding" "sandbox" {
  name      = databricks_catalog.this.name
  # Will this work or is the wokspace id exported as a string?
  workspace = databricks_mws_workspaces.prod.workspace_id
}
```

## Argument Reference

The following arguments are required:

* `name` - Name of Catalog. Change forces creation of a new resource.
* `workspace` - ID of the workspace. Change forces creation of a new resource.

## Import

This resource can be imported by name:

```bash
terraform import databricks_catalog_workspace_binding.this <name>
```
