---
subcategory: "Unity Catalog"
---
# databricks_current_metastore Data Source

Retrieves information about metastore attached to a given workspace.

-> **Note** This is the workspace-level data source.

-> **Note** If you have a fully automated setup with workspaces created by [databricks_mws_workspaces](../resources/mws_workspaces.md) or [azurerm_databricks_workspace](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/databricks_workspace), please make sure to add [depends_on attribute](../guides/troubleshooting.md#data-resources-and-authentication-is-not-configured-errors) to prevent _authentication is not configured for provider_ errors.

## Example Usage

MetastoreSummary response for a metastore attached to the current workspace.

```hcl
data "databricks_current_metastore" "this" {
}

output "some_metastore" {
  value = data.databricks_metastore.this.metastore_info[0]
}
```

## Attribute Reference

This data source exports the following attributes:

* `id` - metastore ID. Will be `no_metastore` if there is no metastore assigned for the current workspace
* `metastore_info` - summary about a metastore attached to the current workspace returned by [Get a metastore summary API](https://docs.databricks.com/api/workspace/metastores/summary). This contains the following attributes (check the API page for up-to-date details):
  * `name` - Name of metastore.
  * `metastore_id` - Metastore ID.
  * `global_metastore_id` - Identifier in form of `<cloud>:<region>:<metastore_id>` for use in Databricks to Databricks Delta Sharing.
  * `region` - (Mandatory for account-level) The region of the metastore.
  * `owner` - Username/group name/sp application_id of the metastore owner.
  * `privilege_model_version` - the version of the privilege model used by the metastore.
  * `storage_root` - Path on cloud storage account, where managed `databricks_table` are stored.
  * `storage_root_credential_id` - ID of a storage credential used for the `storage_root`.
  * `storage_root_credential_name` - Name of a storage credential used for the `storage_root`.
  * `default_data_access_config_id` -  the ID of the default data access configuration.
  * `delta_sharing_scope` - Used to enable delta sharing on the metastore. Valid values: INTERNAL, INTERNAL_AND_EXTERNAL. INTERNAL only allows sharing within the same account, and INTERNAL_AND_EXTERNAL allows cross account sharing and token based sharing.
  * `delta_sharing_recipient_token_lifetime_in_seconds` - the expiration duration in seconds on recipient data access tokens.
  * `delta_sharing_organization_name` - The organization name of a Delta Sharing entity. This field is used for Databricks to Databricks sharing.
  * `created_at` - Timestamp (in milliseconds) when the current metastore was created.
  * `created_by` - the ID of the identity that created the current metastore.
  * `updated_at` - Timestamp (in milliseconds) when the current metastore was updated.
  * `updated_by` - the ID of the identity that updated the current metastore.

## Related Resources

The following resources are used in the same context:

* [databricks_metastore](./metastore.md) to get information for a metastore with a given ID.
* [databricks_metastores](./metastores.md) to get a mapping of name to id of all metastores.
* [databricks_metastore](../resources/metastore.md) to manage Metastores within Unity Catalog.
* [databricks_catalog](../resources/catalog.md) to manage catalogs within Unity Catalog.
