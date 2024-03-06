---
page_title: "Working with Unity Catalog by default"
---

# Working with Unity Catalog by default

Databricks began to [enable new workspaces for Unity Catalog automatically](https://learn.microsoft.com/en-us/azure/databricks/data-governance/unity-catalog/get-started#--automatic-enablement-of-unity-catalog) on November 9, 2023, with a rollout proceeding gradually across accounts. Workspaces that were enabled automatically have the following properties:

- An automatically-provisioned Unity Catalog metastore (unless a Unity Catalog metastore already existed for the workspace region).
- Default privileges for workspace admins, such as the ability to create a catalog or an external database connection.
- No metastore admin (unless an existing Unity Catalog metastore was used and a metastore admin was already assigned).
- No metastore-level storage for managed tables and managed volumes (unless an existing Unity Catalog metastore with metastore-level storage was used).
- A workspace catalog, which, when originally provisioned, is named after your workspace.

This removes the need to manually enable Unity Catalog following [this guide](unity-catalog.md). However, you may need to adjust your Terraform configuration to account for this accordingly

## Removing default privileges for workspace admins

An account admin may decide to remove the default privileges granted to workspace admins, such as the ability to create a catalog or connection. This can be achieved using [databricks_grants](../resources/grants.md), which will override any metastore-level grants not defined in Terraform

```hcl
data "databricks_current_metastore" "this" {
}

resource "databricks_grants" "this" {
  metastore = data.databricks_metastore.this.id
  grant {
    principal  = "Data Engineers"
    privileges = ["CREATE_CATALOG", "CREATE_EXTERNAL_LOCATION"]
  }
  grant {
    principal  = "Data Sharer"
    privileges = ["CREATE_RECIPIENT", "CREATE_SHARE"]
  }
}
```

## Avoiding the automatically-provisioned Unity Catalog metastore

An account admin may pre-create metastores with specific admins in all regions that workspaces will be deployed. This will ensure that new workspaces are automatically assigned to the correct metastore

```hcl
variable "regions" {
  default = ["ap-northeast-1", "eu-west-1"]
}

resource "databricks_metastore" "this" {
  for_each = toset(var.regions)
  name     = "metastore-${each.value}"
  region   = each.value
}
```

## Mandating storage for new catalogs

The automatically-provisioned Unity Catalog metastore does not have metastore-level storage, which means each new catalog has to have a storage location defined

```hcl
# this would fail with "storage location required" error
resource "databricks_catalog" "sandbox" {
  name = "sandbox"
}
```

## Using the workspace catalog

The automatically-provisioned workspace catalog is named after the workspace and initially is bound to that workspace only.

To retrieve this catalog using [databricks_catalogs](../data-sources/catalogs.md)

```hcl
variable workspace_name {}

data "databricks_catalogs" "all" {}

locals {
  default_catalog = [for each in data.databricks_catalogs.all.ids : each if strcontains(each, var.workspace_name)]
}
```

This can then be used to create objects under this catalog, e.g.

```hcl
resource "databricks_schema" "sandbox" {
  catalog_name = local.default_catalog[0]
  name         = "sandbox"
}
```

Or bind this catalog to more workspaces

```hcl
resource "databricks_catalog_workspace_binding" "default_catalog" {
  securable_name = local.default_catalog[0]
  workspace_id   = databricks_mws_workspaces.other.workspace_id
}
```
