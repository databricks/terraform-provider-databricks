---
subcategory: "Unity Catalog"
---
# databricks_metastore_data_access (Resource)

!> **DEPRECATED** This resource is deprecated. Please use [databricks_storage_credential](storage_credential.md) and set it as `storage_root_credential_id` on the [databricks_metastore](metastore.md) resource instead. See the [Unity Catalog API documentation](https://docs.databricks.com/api-explorer/workspace/metastores/create) for more details.

-> This resource can be used with an account or workspace-level provider.

Optionally, each [databricks_metastore](metastore.md) can have a default [databricks_storage_credential](storage_credential.md) defined as `databricks_metastore_data_access`. This will be used by Unity Catalog to access data in the root storage location if defined.

## Migration to databricks_storage_credential

Instead of using `databricks_metastore_data_access`, you should create a [databricks_storage_credential](storage_credential.md) and reference it in your metastore configuration using the `storage_root_credential_id` attribute.

**Old approach (deprecated):**

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

**New approach (although the use of `storage_root` isn't recommended anymore):**

```hcl
resource "databricks_storage_credential" "this" {
  name = aws_iam_role.metastore_data_access.name
  aws_iam_role {
    role_arn = aws_iam_role.metastore_data_access.arn
  }
  comment = "Managed by TF"
}

resource "databricks_metastore" "this" {
  name                        = "primary"
  storage_root                = "s3://${aws_s3_bucket.metastore.id}/metastore"
  owner                       = "uc admins"
  region                      = "us-east-1"
  force_destroy               = true
  storage_root_credential_id  = databricks_storage_credential.this.id
}
```

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

```hcl
import {
  to = databricks_metastore_data_access.this
  id = "<metastore_id>|<name>"
}
```

Alternatively, when using `terraform` version 1.4 or earlier, import using the `terraform import` command:

```bash
terraform import databricks_metastore_data_access.this "<metastore_id>|<name>"
```
