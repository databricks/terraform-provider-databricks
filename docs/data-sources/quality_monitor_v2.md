---
subcategory: "Quality Monitor"
---
# databricks_quality_monitor_v2 Data Source
[![Public Beta](https://img.shields.io/badge/Release_Stage-Public_Beta-orange)](https://docs.databricks.com/aws/en/release-notes/release-types)

~> **Deprecated** This data source is deprecated. Please use `databricks_data_quality_monitor` instead.

This data source can be used to fetch a quality monitors v2.

-> **Note** This data source can only be used with an workspace-level provider!


## Example Usage
~> **Deprecated** This data source is deprecated. Please use `databricks_data_quality_monitor` instead.

Referring to a quality monitor by uc object type (currently only support `schema`) and object id:

```hcl
data "databricks_schema" "this" {
  name = "my_catalog.my_schema"
}
data "databricks_quality_monitor_v2" "this" {
  object_type = "schema"
  object_id = data.databricks_schema.this.schema_info.schema_id
}
```


## Arguments
The following arguments are supported:
* `object_id` (string, required) - The uuid of the request object. For example, schema id
* `object_type` (string, required) - The type of the monitored object. Can be one of the following: schema

## Attributes
The following attributes are exported:
* `anomaly_detection_config` (AnomalyDetectionConfig)
* `object_id` (string) - The uuid of the request object. For example, schema id
* `object_type` (string) - The type of the monitored object. Can be one of the following: schema
* `validity_check_configurations` (list of ValidityCheckConfiguration) - Validity check configurations for anomaly detection

### AnomalyDetectionConfig
* `excluded_table_full_names` (list of string) - List of fully qualified table names to exclude from anomaly detection
* `last_run_id` (string) - Run id of the last run of the workflow
* `latest_run_status` (string) - The status of the last run of the workflow. Possible values are: `ANOMALY_DETECTION_RUN_STATUS_CANCELED`, `ANOMALY_DETECTION_RUN_STATUS_FAILED`, `ANOMALY_DETECTION_RUN_STATUS_JOB_DELETED`, `ANOMALY_DETECTION_RUN_STATUS_PENDING`, `ANOMALY_DETECTION_RUN_STATUS_RUNNING`, `ANOMALY_DETECTION_RUN_STATUS_SUCCESS`, `ANOMALY_DETECTION_RUN_STATUS_UNKNOWN`, `ANOMALY_DETECTION_RUN_STATUS_WORKSPACE_MISMATCH_ERROR`

### PercentNullValidityCheck
* `column_names` (list of string) - List of column names to check for null percentage
* `upper_bound` (number) - Optional upper bound; we should use auto determined bounds for now

### RangeValidityCheck
* `column_names` (list of string) - List of column names to check for range validity
* `lower_bound` (number) - Lower bound for the range
* `upper_bound` (number) - Upper bound for the range

### UniquenessValidityCheck
* `column_names` (list of string) - List of column names to check for uniqueness

### ValidityCheckConfiguration
* `name` (string) - Can be set by system. Does not need to be user facing
* `percent_null_validity_check` (PercentNullValidityCheck)
* `range_validity_check` (RangeValidityCheck)
* `uniqueness_validity_check` (UniquenessValidityCheck)