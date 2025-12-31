---
subcategory: "Unity Catalog"
---
# databricks_metastore Resource

-> This resource can be used with an account or workspace-level provider.

A metastore is the top-level container of objects in Unity Catalog. It stores data assets (tables and views) and the permissions that govern access to them. Databricks account admins can create metastores and assign them to Databricks workspaces in order to control which workloads use each metastore.

Unity Catalog offers a new metastore with built in security and auditing. This is distinct to the metastore used in previous versions of Databricks (based on the Hive Metastore).

A Unity Catalog metastore can be created without a root location & credential to maintain strict separation of storage across catalogs or environments.

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

resource "databricks_metastore_assignment" "this" {
  metastore_id = databricks_metastore.this.id
  workspace_id = local.workspace_id
}
```

For Azure

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

resource "databricks_metastore_assignment" "this" {
  metastore_id = databricks_metastore.this.id
  workspace_id = local.workspace_id
}
```

For GCP

```hcl
resource "databricks_metastore" "this" {
  name          = "primary"
  storage_root  = "gs://${google_storage_bucket.unity_metastore.name}"
  owner         = "uc admins"
  region        = us-east1
  force_destroy = true
}

resource "databricks_metastore_assignment" "this" {
  metastore_id = databricks_metastore.this.id
  workspace_id = local.workspace_id
}
```

## Argument Reference

The following arguments are required:

* `name` - Name of metastore.
* `storage_root` - (Optional) Path on cloud storage account, where managed `databricks_table` are stored.  If the URL contains special characters, such as space, `&`, etc., they should be percent-encoded (space -> `%20`, etc.). Change forces creation of a new resource. If no `storage_root` is defined for the metastore, each catalog must have a `storage_root` defined.  **It's recommended to define `storage_root` on the catalog level.
* `region` - (Mandatory for account-level) The region of the metastore
* `owner` - (Optional) Username/groupname/sp application_id of the metastore owner.
* `delta_sharing_scope` - (Optional) Required along with `delta_sharing_recipient_token_lifetime_in_seconds`. Used to enable delta sharing on the metastore. Valid values: `INTERNAL`, `INTERNAL_AND_EXTERNAL`.  `INTERNAL` only allows sharing within the same account, and `INTERNAL_AND_EXTERNAL` allows cross account sharing and token based sharing.
* `delta_sharing_recipient_token_lifetime_in_seconds` - (Optional) Required along with `delta_sharing_scope`. Used to set expiration duration in seconds on recipient data access tokens.
* `delta_sharing_organization_name` - (Optional) The organization name of a Delta Sharing entity. This field is used for Databricks to Databricks sharing. Once this is set it cannot be removed and can only be modified to another valid value. To delete this value please taint and recreate the resource.
* `external_access_enabled` - (Optional) Whether to allow non-DBR clients to directly access entities under the metastore.
* `force_destroy` - (Optional) Destroy metastore regardless of its contents.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - System-generated ID of this Unity Catalog Metastore. Same as `metastore_id`.
* `cloud` - Cloud vendor of the metastore home shard (e.g., `aws`, `azure`, `gcp`).
* `created_at` - Time at which the metastore was created, in epoch milliseconds.
* `created_by` - Username of metastore creator.
* `default_data_access_config_id` - (Optional) Unique identifier of the metastore's default data access configuration.
* `external_access_enabled` - Whether to allow non-DBR clients to directly access entities under the metastore.
* `global_metastore_id` - Globally unique metastore ID across clouds and regions, of the form `cloud:region:metastore_id`.
* `metastore_id` - Unique identifier of the metastore.
* `privilege_model_version` - Privilege model version of the metastore, of the form `major.minor` (e.g., `1.0`).
* `storage_root_credential_id` - (Optional) UUID of storage credential to access the metastore storage_root.
* `storage_root_credential_name` - Name of the storage credential to access the metastore storage_root.
* `updated_at` - Time at which the metastore was last modified, in epoch milliseconds.
* `updated_by` - Username of user who last modified the metastore.

## Import

This resource can be imported by ID:

```hcl
import {
  to = databricks_metastore.this
  id = "<id>"
}
```

Alternatively, when using `terraform` version 1.4 or earlier, import using the `terraform import` command:

```bash
terraform import databricks_metastore.this <id>
```
