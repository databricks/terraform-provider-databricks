---
subcategory: "Unity Catalog"
---
# databricks_catalog_workspace_binding Resource

If you use workspaces to isolate user data access, you may want to limit catalog access to specific workspaces in your account, also known as workspace-catalog binding

By default, Databricks assigns the catalog to all workspaces attached to the current metastore. By using `databricks_catalog_workspace_binding`, the catalog will be unassigned from all workspaces and only assigned explicitly using this resource.

-> **Note**
  To use this resource the catalog must have its isolation mode set to `ISOLATED` in the [`databricks_catalog`](https://registry.terraform.io/providers/databricks/databricks/latest/docs/resources/catalog#isolation_mode) resource. Alternatively, the isolation mode can be set using the UI or API by following [this guide](https://docs.databricks.com/data-governance/unity-catalog/create-catalogs.html#configuration).

-> **Note**
  If the catalog's isolation mode was set to `ISOLATED` using Terraform then the catalog will have been automatically bound to the workspace it was created from.

## Example Usage

```hcl
resource "databricks_catalog" "sandbox" {
  name           = "sandbox"
  isolation_mode = "ISOLATED"
}

resource "databricks_catalog_workspace_binding" "sandbox" {
  catalog_name = databricks_catalog.sandbox.name
  workspace_id = databricks_mws_workspaces.other.workspace_id
}
```

## Argument Reference

The following arguments are required:

* `workspace_id` - ID of the workspace. Change forces creation of a new resource.
* `securable_name` - Name of securable. Change forces creation of a new resource.
* `securable_type` - Type of securable. Default to `catalog`. Change forces creation of a new resource.
* `binding_type` - Binding mode. Possible values are `BINDING_TYPE_READ_ONLY`, `BINDING_TYPE_READ_WRITE`

## Import

-> **Note** Importing this resource is not currently supported.
