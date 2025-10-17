---
subcategory: "Data Quality Monitoring"
---
# databricks_data_quality_refresh Data Source
[![Private Preview](https://img.shields.io/badge/Release_Stage-Private_Preview-blueviolet)](https://docs.databricks.com/aws/en/release-notes/release-types)

This data source can be used to fetch a data quality refresh on a Unity Catalog table.

The caller must either:
1. be an owner of the table's parent catalog
2. have **USE_CATALOG** on the table's parent catalog and be an owner of the table's parent schema.
3. have the following permissions:
   - **USE_CATALOG** on the table's parent catalog
   - **USE_SCHEMA** on the table's parent schema
   - **SELECT** privilege on the table.

-> **Note** This data source can only be used with a workspace-level provider!

## Example Usage
Getting a data quality refresh by Unity Catalog object type (currently supports `table`) and object id:

```hcl
data "databricks_table" "this" {
  name = "my_catalog.my_schema.my_table"
}
data "databricks_data_quality_refresh" "this" {
  object_type = "table"
  object_id = data.databricks_table.this.id
}
```

## Arguments
The following arguments are supported:
* `object_id` (string, required) - The UUID of the request object. For example, table id
* `object_type` (string, required) - The type of the monitored object. Can be one of the following: `schema`or `table`
* `refresh_id` (integer, required) - Unique id of the refresh operation

## Attributes
The following attributes are exported:
* `end_time_ms` (integer) - Time when the refresh ended (milliseconds since 1/1/1970 UTC)
* `message` (string) - An optional message to give insight into the current state of the refresh (e.g. FAILURE messages)
* `object_id` (string) - The UUID of the request object. For example, table id
* `object_type` (string) - The type of the monitored object. Can be one of the following: `schema`or `table`
* `refresh_id` (integer) - Unique id of the refresh operation
* `start_time_ms` (integer) - Time when the refresh started (milliseconds since 1/1/1970 UTC)
* `state` (string) - The current state of the refresh. Possible values are: `MONITOR_REFRESH_STATE_CANCELED`, `MONITOR_REFRESH_STATE_FAILED`, `MONITOR_REFRESH_STATE_PENDING`, `MONITOR_REFRESH_STATE_RUNNING`, `MONITOR_REFRESH_STATE_SUCCESS`, `MONITOR_REFRESH_STATE_UNKNOWN`
* `trigger` (string) - What triggered the refresh. Possible values are: `MONITOR_REFRESH_TRIGGER_DATA_CHANGE`, `MONITOR_REFRESH_TRIGGER_MANUAL`, `MONITOR_REFRESH_TRIGGER_SCHEDULE`, `MONITOR_REFRESH_TRIGGER_UNKNOWN`