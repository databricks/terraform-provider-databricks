---
subcategory: "Data Quality Monitoring"
---
# databricks_data_quality_refreshes Data Source
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

This data source can be used to fetch the list of data quality refreshes on a Unity Catalog table.

The caller must either:
1. be an owner of the table's parent catalog
2. have **USE_CATALOG** on the table's parent catalog and be an owner of the table's parent schema.
3. have the following permissions:
   - **USE_CATALOG** on the table's parent catalog
   - **USE_SCHEMA** on the table's parent schema
   - **SELECT** privilege on the table.

-> **Note** This data source can only be used with a workspace-level provider!

## Example Usage
Getting a list of all data quality refresh for a given table:

```hcl
data "databricks_table" "this" {
  name = "my_catalog.my_schema.my_table"
}
data "databricks_data_quality_refreshes" "all" {
  object_type = "table"
  object_id = data.databricks_table.this.id
}
```

## Arguments
The following arguments are supported:
* `object_id` (string, required) - The UUID of the request object. It is `schema_id` for `schema`, and `table_id` for `table`.
  
  Find the `schema_id` from either:
  1. The [schema_id](https://docs.databricks.com/api/workspace/schemas/get#schema_id) of the `Schemas` resource.
  2. In [Catalog Explorer](https://docs.databricks.com/aws/en/catalog-explorer/) > select the `schema` > go to the `Details` tab > the `Schema ID` field.
  
  Find the `table_id` from either:
  1. The [table_id](https://docs.databricks.com/api/workspace/tables/get#table_id) of the `Tables` resource.
  2. In [Catalog Explorer](https://docs.databricks.com/aws/en/catalog-explorer/) > select the `table` > go to the `Details` tab > the `Table ID` field
* `object_type` (string, required) - The type of the monitored object. Can be one of the following: `schema` or `table`
* `page_size` (integer, optional)


## Attributes
This data source exports a single attribute, `refreshes`. It is a list of resources, each with the following attributes:
* `end_time_ms` (integer) - Time when the refresh ended (milliseconds since 1/1/1970 UTC)
* `message` (string) - An optional message to give insight into the current state of the refresh (e.g. FAILURE messages)
* `object_id` (string) - The UUID of the request object. It is `schema_id` for `schema`, and `table_id` for `table`.
  
  Find the `schema_id` from either:
  1. The [schema_id](https://docs.databricks.com/api/workspace/schemas/get#schema_id) of the `Schemas` resource.
  2. In [Catalog Explorer](https://docs.databricks.com/aws/en/catalog-explorer/) > select the `schema` > go to the `Details` tab > the `Schema ID` field.
  
  Find the `table_id` from either:
  1. The [table_id](https://docs.databricks.com/api/workspace/tables/get#table_id) of the `Tables` resource.
  2. In [Catalog Explorer](https://docs.databricks.com/aws/en/catalog-explorer/) > select the `table` > go to the `Details` tab > the `Table ID` field
* `object_type` (string) - The type of the monitored object. Can be one of the following: `schema`or `table`
* `refresh_id` (integer) - Unique id of the refresh operation
* `start_time_ms` (integer) - Time when the refresh started (milliseconds since 1/1/1970 UTC)
* `state` (string) - The current state of the refresh. Possible values are: `MONITOR_REFRESH_STATE_CANCELED`, `MONITOR_REFRESH_STATE_FAILED`, `MONITOR_REFRESH_STATE_PENDING`, `MONITOR_REFRESH_STATE_RUNNING`, `MONITOR_REFRESH_STATE_SUCCESS`, `MONITOR_REFRESH_STATE_UNKNOWN`
* `trigger` (string) - What triggered the refresh. Possible values are: `MONITOR_REFRESH_TRIGGER_DATA_CHANGE`, `MONITOR_REFRESH_TRIGGER_MANUAL`, `MONITOR_REFRESH_TRIGGER_SCHEDULE`, `MONITOR_REFRESH_TRIGGER_UNKNOWN`