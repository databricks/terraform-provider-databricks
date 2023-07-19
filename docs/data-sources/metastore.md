---
subcategory: "Unity Catalog"
---
# databricks_metastore Data Source

Retrieves information about metastore for a given id of [databricks_metastore](../resources/metastore.md) object, that was created by Terraform or manually, so that special handling could be applied.

-> **Note** If you have a fully automated setup with workspaces created by [databricks_mws_workspaces](../resources/mws_workspaces.md) or [azurerm_databricks_workspace](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/databricks_workspace), please make sure to add [depends_on attribute](../index.md#data-resources-and-authentication-is-not-configured-errors) in order to prevent _authentication is not configured for provider_ errors.

## Example Usage

MetastoreInfo response for a given metastore id

```hcl

resource "databricks_metastore" "this" {
  provider      = databricks.workspace
  name          = "primary"
  storage_root  = "s3://${aws_s3_bucket.metastore.id}/metastore"
  owner         = var.unity_admin_group
  force_destroy = true
}

data "databricks_metastore" "this" {
  id = databricks_metastore.this.id
}

output "some_metastore" {
  value = data.databricks_metastore.this.metastore_info[0]
}
```

## Argument Reference
* `metastore_id` - Id of the metastore to be fetched

## Attribute Reference

This data source exports the following attributes:
* `metastore_info` - MetastoreInfo object for a [databricks_metastore](../resources/metastore.md). This contains the following attributes:
  * `name` - Name of metastore.
  * `storage_root` - Path on cloud storage account, where managed `databricks_table` are stored. Change forces creation of a new resource.
  * `owner` - Username/groupname/sp application_id of the metastore owner.
  * `delta_sharing_scope` - Used to enable delta sharing on the metastore. Valid values: INTERNAL, INTERNAL_AND_EXTERNAL.
  * `delta_sharing_recipient_token_lifetime_in_seconds` - Used to set expiration duration in seconds on recipient data access tokens.
  * `delta_sharing_organization_name` - The organization name of a Delta Sharing entity. This field is used for Databricks to Databricks sharing. 


## Related Resources

The following resources are used in the same context:
* [databricks_metastores](./metastores.md) to get mapping of name to id of all metastores.
* [databricks_metastore](../resources/metastore.md) to manage Metastores within Unity Catalog.
* [databricks_catalog](../resources/catalog.md) to manage catalogs within Unity Catalog.