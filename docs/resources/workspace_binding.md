---
subcategory: "Unity Catalog"
---
# databricks_workspace_binding Resource

-> **Note** This resource could be only used with workspace-level provider!

If you use workspaces to isolate user data access, you may want to limit access to catalog, external locations or storage credentials from specific workspaces in your account, also known as workspace binding

By default, Databricks assigns the securable to all workspaces attached to the current metastore. By using `databricks_workspace_binding`, the securable will be unassigned from all workspaces and only assigned explicitly using this resource.

-> **Note**
  To use this resource the securable must have its isolation mode set to `ISOLATED` in the. Alternatively, the isolation mode can be set using the UI or API by following [this guide](https://docs.databricks.com/data-governance/unity-catalog/create-catalogs.html#configuration), [this guide](https://docs.databricks.com/en/connect/unity-catalog/external-locations.html#workspace-binding) or [this guide](https://docs.databricks.com/en/connect/unity-catalog/storage-credentials.html#optional-assign-a-storage-credential-to-specific-workspaces).

-> **Note**
  If the securable's isolation mode was set to `ISOLATED` using Terraform then the securable will have been automatically bound to the workspace it was created from.

## Example Usage

```hcl
resource "databricks_catalog" "sandbox" {
  name           = "sandbox"
  isolation_mode = "ISOLATED"
}

resource "databricks_catalog_workspace_binding" "sandbox" {
  securable_name = databricks_catalog.sandbox.name
  workspace_id   = databricks_mws_workspaces.other.workspace_id
}
```

## Argument Reference

The following arguments are required:

* `workspace_id` - ID of the workspace. Change forces creation of a new resource.
* `securable_name` - Name of securable. Change forces creation of a new resource. Can be `catalog`, `external-locations` or `storage-credentials`.
* `securable_type` - Type of securable. Default to `catalog`. Change forces creation of a new resource.
* `binding_type` - (Optional) Binding mode. Default to `BINDING_TYPE_READ_WRITE`. For `catalog`, possible values are `BINDING_TYPE_READ_ONLY`, `BINDING_TYPE_READ_WRITE`. For `external-locations` or `storage-credentials`, no binding mode needs to be specified

## Import

This resource can be imported by using combination of workspace ID, securable type and name:

```sh
terraform import databricks_catalog_workspace_binding.this "<workspace_id>|<securable_type>|<securable_name>"
```
