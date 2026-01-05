---
subcategory: "Unity Catalog"
---
# databricks_metastore Data Source

Retrieves information about metastore for a given id of [databricks_metastore](../resources/metastore.md) object, that was created by Terraform or manually, so that special handling could be applied.

-> This data source can only be used with an account-level provider!

## Example Usage

MetastoreInfo response for a given metastore id

```hcl
resource "aws_s3_bucket" "metastore" {
  bucket        = "${var.prefix}-metastore"
  force_destroy = true
}

resource "databricks_metastore" "this" {
  provider      = databricks.workspace
  name          = "primary"
  storage_root  = "s3://${aws_s3_bucket.metastore.id}/metastore"
  owner         = var.unity_admin_group
  force_destroy = true
}

data "databricks_metastore" "this" {
  metastore_id = databricks_metastore.this.id
}

output "some_metastore" {
  value = data.databricks_metastore.this.metastore_info[0]
}
```

## Argument Reference

Provide one of the arguments to get information about a metastore:

* `metastore_id` - ID of the metastore
* `name` - Name of the metastore
* `region` - Region of the metastore

## Attribute Reference

This data source exports the following attributes:

* `id` - ID of the metastore
* `metastore_info` - MetastoreInfo object for a [databricks_metastore](../resources/metastore.md). This contains the following attributes:
  * `cloud` - Cloud vendor of the metastore home shard (e.g., `aws`, `azure`, `gcp`).
  * `created_at` - Time at which the metastore was created, in epoch milliseconds.
  * `created_by` - Username of metastore creator.
  * `default_data_access_config_id` - Unique identifier of the metastore's default data access configuration.
  * `delta_sharing_organization_name` - The organization name of a Delta Sharing entity. This field is used for Databricks to Databricks sharing.
  * `delta_sharing_recipient_token_lifetime_in_seconds` - Used to set expiration duration in seconds on recipient data access tokens.
  * `delta_sharing_scope` - Used to enable delta sharing on the metastore. Valid values: INTERNAL, INTERNAL_AND_EXTERNAL. INTERNAL only allows sharing within the same account, and INTERNAL_AND_EXTERNAL allows cross account sharing and token based sharing.
  * `external_access_enabled` - Whether to allow non-DBR clients to directly access entities under the metastore.
  * `global_metastore_id` - Globally unique metastore ID across clouds and regions, of the form `cloud:region:metastore_id`.
  * `metastore_id` - Unique identifier of the metastore.
  * `name` - Name of metastore.
  * `owner` - Username/groupname/sp application_id of the metastore owner.
  * `privilege_model_version` - Privilege model version of the metastore, of the form `major.minor` (e.g., `1.0`).
  * `region` - Cloud region which the metastore serves (e.g., `us-west-2`, `westus`).
  * `storage_root_credential_id` - UUID of storage credential to access the metastore storage_root.
  * `storage_root_credential_name` - Name of the storage credential to access the metastore storage_root.
  * `storage_root` - Path on cloud storage account, where managed `databricks_table` are stored.
  * `updated_at` - Time at which the metastore was last modified, in epoch milliseconds.
  * `updated_by` - Username of user who last modified the metastore.

## Related Resources

The following resources are used in the same context:

* [databricks_metastores](./metastores.md) to get mapping of name to id of all metastores.
* [databricks_metastore](../resources/metastore.md) to manage Metastores within Unity Catalog.
* [databricks_catalog](../resources/catalog.md) to manage catalogs within Unity Catalog.
