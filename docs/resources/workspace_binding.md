---
subcategory: "Unity Catalog"
---
# databricks_workspace_binding Resource

If you use workspaces to isolate user data access, you may want to limit access to catalog, external locations or storage credentials from specific workspaces in your account, also known as workspace binding

-> This resource can only be used with a workspace-level provider!

By default, Databricks assigns the securable to all workspaces attached to the current metastore. By using `databricks_workspace_binding`, the securable will be unassigned from all workspaces and only assigned explicitly using this resource.

-> To use this resource the securable must have its isolation mode set to `ISOLATED` (for [databricks_catalog](catalog.md)) or `ISOLATION_MODE_ISOLATED` (for  (for [databricks_external_location](external_location.md), [databricks_storage_credential](storage_credential.md) or [databricks_credential](credential.md)) for the `isolation_mode` attribute. Alternatively, the isolation mode can be set using the UI or API by following [this guide](https://docs.databricks.com/data-governance/unity-catalog/create-catalogs.html#configuration), [this guide](https://docs.databricks.com/en/connect/unity-catalog/external-locations.html#workspace-binding) or [this guide](https://docs.databricks.com/en/connect/unity-catalog/storage-credentials.html#optional-assign-a-storage-credential-to-specific-workspaces).

-> If the securable's isolation mode was set to `ISOLATED` using Terraform then the securable will have been automatically bound to the workspace it was created from.

## Example Usage

```hcl
resource "databricks_catalog" "sandbox" {
  name           = "sandbox"
  isolation_mode = "ISOLATED"
}

resource "databricks_workspace_binding" "sandbox" {
  securable_name = databricks_catalog.sandbox.name
  workspace_id   = databricks_mws_workspaces.other.workspace_id
}
```

## Argument Reference

The following arguments are required:

* `workspace_id` - ID of the workspace. Change forces creation of a new resource.
* `securable_name` - Name of securable. Change forces creation of a new resource.
* `securable_type` - Type of securable. Can be `catalog`, `external_location`, `storage_credential` or `credential`. Default to `catalog`. Change forces creation of a new resource.
* `binding_type` - (Optional) Binding mode. Default to `BINDING_TYPE_READ_WRITE`. Possible values are `BINDING_TYPE_READ_ONLY`, `BINDING_TYPE_READ_WRITE`.

## Migration from databricks_catalog_workspace_binding

You can migrate from the deprecated `databricks_catalog_workspace_binding` to `databricks_workspace_binding` without re-binding catalog.

### For Terraform version >= 1.7.0

Terraform 1.7 introduced the [removed](https://developer.hashicorp.com/terraform/language/resources/syntax#removing-resources) block in addition to the [import](https://developer.hashicorp.com/terraform/language/import) block introduced in Terraform 1.5. Together they make import and removal of resources easier, avoiding manual execution of `terraform import` and `terraform state rm` commands.

So with Terraform 1.7+, the migration looks as the following:

* Remove the `databricks_catalog_workspace_binding` resource and replace it with the `databricks_workspace_binding`.
* Add `import` and `removed` blocks like this:

```hcl
locals {
  workspace_id = 1234567890
}

removed {
  from = databricks_catalog_workspace_binding.sandbox
  lifecycle {
    destroy = false
  }
}

resource "databricks_workspace_binding" "sandbox" {
  securable_name = databricks_catalog.sandbox.name
  workspace_id   = local.workspace_id
}

import {
  to = databricks_workspace_binding.sandbox
  id = "${local.workspace_id}|catalog|${databricks_catalog.sandbox.name}"
}
```

* Run the `terraform plan` command to check possible changes, such as value type change, etc.
* Run the `terraform apply` command to apply changes.
* Remove the `import` and `removed` blocks from the code.

### For Terraform version < 1.7.0

* Remove the `databricks_catalog_workspace_binding` resource and and replace it with the `databricks_workspace_binding`.
* Remove the old resource from the state with the `terraform state rm databricks_catalog_workspace_binding.sandbox` command.
* Import new resource with the `terraform import databricks_workspace_binding.sandbox "<workspace_id>|<securable_type>|<securable_name>"` command.
* Run the `terraform plan` command to check possible changes, such as value type change, etc.

## Import

This resource can be imported by using combination of workspace ID, securable type and name:

```hcl
import {
  to = databricks_workspace_binding.this
  id = "<workspace_id>|<securable_type>|<securable_name>"
}
```

Alternatively, when using `terraform` version 1.4 or earlier, import using the `terraform import` command:

```bash
terraform import databricks_workspace_binding.this "<workspace_id>|<securable_type>|<securable_name>"
```
