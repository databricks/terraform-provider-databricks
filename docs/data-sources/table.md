---
subcategory: "Unity Catalog"
---
# databricks_table Data Source

-> **Note** If you have a fully automated setup with workspaces created by [databricks_mws_workspaces](../resources/mws_workspaces.md) or [azurerm_databricks_workspace](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/databricks_workspace), please make sure to add [depends_on attribute](../guides/troubleshooting.md#data-resources-and-authentication-is-not-configured-errors) in order to prevent _default auth: cannot configure default credentials_ errors.

Retrieves a single table from Unity Catalog, including all fields returned by the Databricks SDK

## Example Usage
```hcl
data "databricks_table" "thing1" {
  full_name = "sandbox.things.thing1"
}
```

## Argument Reference

* `full_name` - (Required) Fully qualified table name - *`catalog`.`schema`.`table`*

## Attribute Reference

This data source exports the same attributes as the Databricks SDK [TableInfo](https://pkg.go.dev/github.com/databricks/databricks-sdk-go/service/catalog#TableInfo):

Notable fields are:
* `access_point` - The AWS access point to use when accesing s3 for this external locations
* `catalog_name` - Name of parent catalog
* `columns` - The array of __ColumnInfo__ definitions of the table's columns
  * `comment` - User-provided free-form text description
  * `mask` - Mask function
  * `name` - Name of column
  * `nullable` - Whether the column is nullable
  * `partition_index` - Partition order or null
  * `position` - Column order
  * `type_interval_type` - Interval type if applicable
  * `type_json` - JSON with keys (name, type, date, nullable, metadata)
  * `type_name` - String name of type
  * `type_precision` - Type precision
  * `type_scale` - Scale of type
  * `type_text` - Text type name (lowercase)
* `comment` - User-provided free-form text description
* `created_at` - Time at which this table was created, in epoch milliseconds
* `created_by` - Username of table creator
* `data_access_configuration_id` - Unique ID of the Data Access Configuration to use with the table data
* `data_source_format` - Data source format
* `deleted_at` - Time at which this table was deleted, in epoch milliseconds. Field is omitted if table is not deleted
* `delta_runtime_properties_kvpairs` - Information pertaining to current state of the delta table
* `effective_predictive_optimization_flag`
* `enable_predictive_optimization` - Whether predictive optimization should be enabled for this object and objects under it
* `encryption_details` - Encryption options that apply to clients connecting to cloud storage
* `metastore_id` - Unique identifier of parent metastore
* `name` - Name of table, relative to parent schema
* `owner` - Username of current owner of table
* `pipeline_id` - The pipeline ID of the table. Applicable for tables created by pipelines (Materialized View, Streaming Table, etc.)
* `properties` - A map of key-value properties attached to the securable
* `row_filter` - JSON configuration of a row filter
* `schema_name` - Name of parent schema relative to its parent catalog
* `sql_path` - List of schemes whose objects can be referenced without qualification
* `storage_credential_name` - Name of the storage credential, when a storage credential is configured for use with this table
* `storage_location` - Storage root URL for table (for **MANAGED**, **EXTERNAL** tables)
* `table_constraints` - List of table constraints
* `table_id` - Name of table, relative to parent schema
* `table_type` - Table or View
* `updated_at` - Time at which this table was last modified, in epoch milliseconds
* `updated_by` - Username of user who last modified the table
* `view_definition` - View definition SQL (when __table_type__ is **VIEW**, **MATERIALIZED_VIEW**, or **STREAMING_TABLE**)
* `view_dependencies` - View dependencies (when table_type == **VIEW** or **MATERIALIZED_VIEW**, **STREAMING_TABLE**) - when DependencyList is None, the dependency is not provided; - when DependencyList is an empty list, the dependency is provided but is empty; - when DependencyList is not an empty list, dependencies are provided and recorded.

## Related Resources

The following resources are used in the same context:

* [databricks_schema](../resources/schema.md) to manage schemas within Unity Catalog.
* [databricks_catalog](../resources/catalog.md) to manage catalogs within Unity Catalog.

