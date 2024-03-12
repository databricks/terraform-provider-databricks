---
subcategory: "Unity Catalog"
---
# databricks_metastore_data_access (Resource)

-> **Note** This resource could be used with account or workspace-level provider.

Optionally, each [databricks_metastore](metastore.md) can have a default [databricks_storage_credential](storage_credential.md) defined as `databricks_metastore_data_access`. This will be used by Unity Catalog to access data in the root storage location if defined.

## Example Usage

For AWS

```hcl
resource "databricks_metastore" "this" {
  name          = "primary"
  storage_root  = "s3://${aws_s3_bucket.metastore.id}/metastore"
  owner         = "uc admins"
  region        = "us-east-1"
  force_destroy = true
}

resource "databricks_metastore_data_access" "this" {
  metastore_id = databricks_metastore.this.id
  name         = aws_iam_role.metastore_data_access.name
  aws_iam_role {
    role_arn = aws_iam_role.metastore_data_access.arn
  }
  is_default = true
}
```

For Azure using managed identity as credential (recommended)

```hcl
resource "databricks_metastore" "this" {
  name = "primary"
  storage_root = format("abfss://%s@%s.dfs.core.windows.net/",
    azurerm_storage_container.unity_catalog.name,
  azurerm_storage_account.unity_catalog.name)
  owner         = "uc admins"
  region        = "eastus"
  force_destroy = true
}

resource "databricks_metastore_data_access" "this" {
  metastore_id = databricks_metastore.this.id
  name         = "mi_dac"
  azure_managed_identity {
    access_connector_id = var.access_connector_id
  }
  is_default = true
}
```

## Argument Reference

The arguments are the same as of [databricks_storage_credential](storage_credential.md). Additionally

* `is_default` -  whether to set this credential as the default for the metastore. In practice, this should always be true.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of this data access configuration in form of `<metastore_id>|<name>`.

## Import

This resource can be imported by combination of metastore id and the data access name.

```bash
terraform import databricks_metastore_data_access.this '<metastore_id>|<name>'
```
